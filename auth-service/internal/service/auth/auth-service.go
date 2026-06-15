package service

import (
	"auth-service/pkg/jwt"
	"context"
)

type AuthService interface {
	Register(ctx context.Context, register *RegisterRequest) (*AuthResponse, error)
	Login(ctx context.Context, login *LoginRequest) (*AuthResponse, error)
	ChangePassword(ctx context.Context, change *ChangePasswordRequest) error
	RecoverPassword(ctx context.Context, recover *RecoverPasswordRequest) (*RecoverPasswordResponse, error)
	ValidateToken(ctx context.Context, validate *ValidateTokenRequest) (*jwt.Claims, error)
}

type RegisterRequest struct {
	Email     string
	FirstName string
	LastName  string
	Phone     string
	Password  string
}

type LoginRequest struct {
	Email    string
	Password string
}

type ChangePasswordRequest struct {
	UserID      string
	OldPassword string
	NewPassword string
}

type RecoverPasswordRequest struct {
	Email string
}

type ValidateTokenRequest struct {
	Token string
}

type RecoverPasswordResponse struct {
	ImitationNewPassword string
}

type AuthResponse struct {
	Token  string
	UserId string
}
