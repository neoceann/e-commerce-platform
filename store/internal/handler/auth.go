package handler

import (
    "encoding/json"
    "net/http"
    
    "store/internal/grpc/pb"
    "store/internal/grpc"
)

type AuthHandler struct {
    authClient *grpc.AuthClient
}

func NewAuthHandler(authClient *grpc.AuthClient) *AuthHandler {
    return &AuthHandler{
        authClient: authClient,
    }
}

type RegisterRequest struct {
    Email     string `json:"email"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Phone     string `json:"phone"`
    Password  string `json:"password"`
}

type AuthRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type AuthResponse struct {
    Token  string `json:"token"`
    UserID string `json:"user_id"`
    Error  string `json:"error,omitempty"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeError(w, http.StatusBadRequest, "invalid request")
        return
    }
    
    resp, err := h.authClient.GetClient().Register(r.Context(), &pb.RegisterRequest{
        Email:     req.Email,
        FirstName: req.FirstName,
        LastName:  req.LastName,
        Phone:     req.Phone,
        Password:  req.Password,
    })
    
    if err != nil {
        writeError(w, http.StatusInternalServerError, err.Error())
        return
    }
    
    if !resp.Success {
        writeError(w, http.StatusBadRequest, resp.Message)
        return
    }
    
    writeJSON(w, http.StatusCreated, AuthResponse{
        Token:  resp.Token,
        UserID: resp.UserId,
    })
}

func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
    var req AuthRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeError(w, http.StatusBadRequest, "invalid request")
        return
    }
    
    resp, err := h.authClient.GetClient().Login(r.Context(), &pb.LoginRequest{
        Email:    req.Email,
        Password: req.Password,
    })
    
    if err != nil {
        writeError(w, http.StatusInternalServerError, err.Error())
        return
    }
    
    if !resp.Success {
        writeError(w, http.StatusUnauthorized, resp.Message)
        return
    }
    
    writeJSON(w, http.StatusOK, AuthResponse{
        Token:  resp.Token,
        UserID: resp.UserId,
    })
}