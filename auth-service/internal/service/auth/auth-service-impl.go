package service

import (
    "context"
    "fmt"
    
    "auth-service/internal/repository/db"
    "auth-service/pkg/jwt"
    "auth-service/pkg/password"
)

type authService struct {
    userRepo *db.Queries
    jwtService   jwt.JWTService
    passwordService   password.PasswordService
}

func NewAuthService(userRepo *db.Queries, j jwt.JWTService, p password.PasswordService) AuthService {
    return &authService{
        userRepo: userRepo,
        jwtService:   j,
        passwordService:   p,
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
	
}
	
func (a *authService) ChangePassword(ctx context.Context, change *ChangePasswordRequest) error {
	
}
	
func (a *authService) RecoverPassword(ctx context.Context, recover *RecoverPasswordRequest) (*RecoverPasswordResponse, error) {
	
}
	
func (a *authService) ValidateToken(ctx context.Context, validate *ValidateTokenRequest) (*jwt.Claims, error) {

}
