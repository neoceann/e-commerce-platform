package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"auth-service/internal/repository/db"
	"auth-service/pkg/jwt"
	"auth-service/pkg/password"

	"github.com/google/uuid"
)

type authService struct {
	userRepo        *db.Queries
	jwtService      jwt.JWTService
	passwordService password.PasswordService
}

func NewAuthService(userRepo *db.Queries, j jwt.JWTService, p password.PasswordService) AuthService {
	return &authService{
		userRepo:        userRepo,
		jwtService:      j,
		passwordService: p,
	}
}

func (a *authService) Register(ctx context.Context, register *RegisterRequest) (*AuthResponse, error) {
	exists, err := a.userRepo.UserExists(ctx, register.Email)

	if err != nil {
		return nil, fmt.Errorf("check user error: %w", err)
	}

	if exists {
		return nil, ErrUserAlreadyExists
	}

	hashpwd, err := a.passwordService.Hash(register.Password)

	if err != nil {
		return nil, fmt.Errorf("hash password error: %w", err)
	}

	userParams := RegisterInfoFromServiceToDB(register)
	userParams.PasswordHash = hashpwd
	userParams.UserRole = "user"

	user, err := a.userRepo.CreateUser(ctx, *userParams)

	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	token, err := a.jwtService.GenerateToken(user.ID.String(), user.Email, user.UserRole)

	if err != nil {
		return nil, fmt.Errorf("generate jwt token error: %w", err)
	}

	return &AuthResponse{Token: token, UserId: user.ID.String()}, nil
}

func (a *authService) Login(ctx context.Context, login *LoginRequest) (*AuthResponse, error) {
	user, err := a.userRepo.GetUserByEmail(ctx, login.Email)

	if err != nil {
		return nil, ErrCredentialsInvalid
	}

	if !a.passwordService.Verify(user.PasswordHash, login.Password) {
		return nil, ErrCredentialsInvalid
	}

	token, err := a.jwtService.GenerateToken(user.ID.String(), user.Email, user.UserRole)

	if err != nil {
		return nil, fmt.Errorf("generate jwt token error: %w", err)
	}

	return &AuthResponse{Token: token, UserId: user.ID.String()}, nil
}

func (a *authService) ChangePassword(ctx context.Context, change *ChangePasswordRequest) error {
	uid, err := uuid.Parse(change.UserID)
	if err != nil {
		return fmt.Errorf("parse uuid error: %w", err)
	}
	user, err := a.userRepo.GetUserByID(ctx, uid)

	if err != nil {
		return ErrUserNotFound
	}

	if !a.passwordService.Verify(user.PasswordHash, change.OldPassword) {
		return ErrInvalidOldPwd
	}

	newPwd, err := a.passwordService.Hash(change.NewPassword)

	if err != nil {
		return fmt.Errorf("hash passoword error: %w", err)
	}

	return a.userRepo.UpdatePassword(ctx, db.UpdatePasswordParams{ID: user.ID, PasswordHash: newPwd})
}

func (a *authService) RecoverPassword(ctx context.Context, recover *RecoverPasswordRequest) (*RecoverPasswordResponse, error) {
	user, err := a.userRepo.GetUserByEmail(ctx, recover.Email)

	if err != nil {
		return nil, ErrUserNotFound
	}

	temp, err := generateRandomPassword(10)

	if err != nil {
		return nil, fmt.Errorf("generate temp password error: %w", err)
	}

	hashed, err := a.passwordService.Hash(temp)

	if err != nil {
		return nil, fmt.Errorf("hash temp password error: %w", err)
	}

	err = a.userRepo.UpdatePassword(ctx, db.UpdatePasswordParams{ID: user.ID, PasswordHash: hashed})

	if err != nil {
		return nil, fmt.Errorf("update temp password error: %w", err)
	}

	log.Printf("PASSWORD RECOVERING FOR USER: %s. New temporary password: %s", user.Email, temp)

	return &RecoverPasswordResponse{ImitationNewPassword: temp}, nil

}

func generateRandomPassword(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (a *authService) ValidateToken(ctx context.Context, validate *ValidateTokenRequest) (*jwt.Claims, error) {
	claims, err := a.jwtService.ValidateToken(validate.Token)

	if err != nil {
		return nil, ErrTokenInvalid
	}

	return claims, nil
}
