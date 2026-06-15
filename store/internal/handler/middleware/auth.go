package middleware

import (
	"context"
	"net/http"

	"store/internal/grpc"
	"store/internal/grpc/pb"
)

type AuthMiddleware struct {
	authClient *grpc.AuthClient
}

func NewAuthMiddleware(authClient *grpc.AuthClient) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: authClient,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r)
		if token == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		resp, err := m.authClient.GetClient().ValidateToken(r.Context(), &pb.ValidateTokenRequest{
			Token: token,
		})

		if err != nil || !resp.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", resp.UserId)
		ctx = context.WithValue(ctx, "user_email", resp.Email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	return authHeader
}

func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value("user_id").(string)
	return userID, ok
}

func GetUserEmail(ctx context.Context) (string, bool) {
	email, ok := ctx.Value("user_email").(string)
	return email, ok
}
