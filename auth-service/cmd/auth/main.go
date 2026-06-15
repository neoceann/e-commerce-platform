package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"auth-service/internal/config"
	server "auth-service/internal/grpc/auth-server"
	pb "auth-service/internal/grpc/pb"
	"auth-service/internal/repository/db"
	"auth-service/internal/service/auth"
	"auth-service/pkg/jwt"
	"auth-service/pkg/password"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"

	 "google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	if err := pool.Ping(ctx); err != nil {
    log.Fatalf("Failed to ping database: %v", err)
	}
	defer pool.Close()

	queries := db.New(pool)
	jwtSvc := jwt.NewJWTService(cfg.JWTSecret, cfg.JWTExpirationToTime())
	pwdSvc := password.NewPasswordService(cfg.BcryptCost)
	authService := service.NewAuthService(queries, jwtSvc, pwdSvc)

	authServer := server.NewAuthServer(authService)

	lis, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authServer)

	reflection.Register(grpcServer)

	go func() {
		log.Printf("gRPC server listening on: %s", cfg.GrpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}
