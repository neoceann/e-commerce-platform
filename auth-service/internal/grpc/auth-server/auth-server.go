package grpc

import (
	"auth-service/internal/grpc/pb"
	"auth-service/internal/service/auth"
	"context"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	authService service.AuthService
}

func NewAuthServer(authService service.AuthService) *AuthServer {
	return &AuthServer{authService: authService}
}

func (a *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	response, err := a.authService.Register(ctx, &service.RegisterRequest{Email: req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Password:  req.Password})

	if err != nil {
		return &pb.AuthResponse{Success: false, Message: err.Error()}, nil
	}

	return &pb.AuthResponse{Success: true,
		Token:   response.Token,
		UserId:  response.UserId,
		Message: "Registration successful"}, nil
}

func (a *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	response, err := a.authService.Login(ctx, &service.LoginRequest{Email: req.Email, Password: req.Password})

	if err != nil {
		return &pb.AuthResponse{Success: false, Message: err.Error()}, nil
	}

	return &pb.AuthResponse{Success: true, Token: response.Token, UserId: response.UserId, Message: "Login successful"}, nil
}

func (a *AuthServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	claims, err := a.authService.ValidateToken(ctx, &service.ValidateTokenRequest{Token: req.Token})

	if err != nil {
		return &pb.ChangePasswordResponse{Success: false, Message: err.Error()}, nil
	}

	err = a.authService.ChangePassword(ctx, &service.ChangePasswordRequest{UserID: claims.ID, OldPassword: req.OldPassword, NewPassword: req.NewPassword})

	if err != nil {
		return &pb.ChangePasswordResponse{Success: false, Message: err.Error()}, nil
	}

	return &pb.ChangePasswordResponse{Success: true, Message: "Password successfully changed"}, nil
}

func (a *AuthServer) RecoverPassword(ctx context.Context, req *pb.RecoverPasswordRequest) (*pb.RecoverPasswordResponse, error) {
	_, err := a.authService.RecoverPassword(ctx, &service.RecoverPasswordRequest{Email: req.Email})

	if err != nil {
		return &pb.RecoverPasswordResponse{Success: false, Message: err.Error()}, nil
	}

	return &pb.RecoverPasswordResponse{Success: true, Message: "Password successfully recovered"}, nil

}

func (a *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	claims, err := a.authService.ValidateToken(ctx, &service.ValidateTokenRequest{Token: req.Token})

	if err != nil {
		return &pb.ValidateTokenResponse{Valid: false, Error: err.Error()}, nil
	}

	return &pb.ValidateTokenResponse{Valid: true, UserId: claims.UserID, Email: claims.Email, Role: claims.Role}, nil
}
