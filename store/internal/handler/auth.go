package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"store/internal/grpc"
	"store/internal/grpc/pb"
	"store/internal/handler/middleware"
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
	Token   string `json:"token"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

type RecoverPasswordRequest struct {
	Email string `json:"email"`
}

type RecoverPasswordResponse struct {
	Message string
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old"`
	NewPassword string `json:"new"`
}

// @Summary      Смена пароля
// @Tags	     Управление учетной записью
// @Security     BearerAuth
// @Param        request body ChangePasswordRequest true "Запрос на смену пароля"
// @Success      200 {object} string
// @Failure      400 {object} string "Неверный запрос"
// @Failure      500 {object} string "Внутренняя ошибка сервера"
// @Router       /change [post]
func (h *AuthHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var req ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request")
		return
	}

	userid, ok := middleware.GetUserID(r.Context())

	if !ok {
		writeError(w, http.StatusBadRequest, "cant get userId")
	}

	log.Printf("user id: %s", userid)

	resp, err := h.authClient.GetClient().ChangePassword(r.Context(), &pb.ChangePasswordRequest{UserId: userid, OldPassword: req.OldPassword, NewPassword: req.NewPassword})

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !resp.Success {
		writeError(w, http.StatusBadRequest, resp.Message)
		return
	}

	writeJSON(w, http.StatusOK, resp.Message)
}

// @Summary      Восстановление пароля
// @Tags	     Управление учетной записью
// @Param        request body RecoverPasswordRequest true "Запрос на восстановление пароля"
// @Success      200 {object} RecoverPasswordResponse
// @Failure      400 {object} string "Неверный запрос"
// @Failure      500 {object} string "Внутренняя ошибка сервера"
// @Router       /recover [post]
func (h *AuthHandler) Recover(w http.ResponseWriter, r *http.Request) {
	var req RecoverPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request")
		return
	}

	resp, err := h.authClient.GetClient().RecoverPassword(r.Context(), &pb.RecoverPasswordRequest{Email: req.Email})

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !resp.Success {
		writeError(w, http.StatusBadRequest, resp.Message)
		return
	}

	writeJSON(w, http.StatusOK, RecoverPasswordResponse{Message: resp.Message})
}

// @Summary      Регистрация нового пользователя
// @Tags	     Управление учетной записью
// @Param        request body RegisterRequest true "Данные для регистрации"
// @Success      201 {object} string
// @Failure      400 {object} string "Неверный запрос"
// @Failure      500 {object} AuthResponse "Внутренняя ошибка сервера"
// @Router       /register [post]
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
		Message: resp.Message,
		Token:   resp.Token,
		UserID:  resp.UserId,
	})
}

// @Summary      Вход зарегистрированного пользователя
// @Tags	     Управление учетной записью
// @Param        request body AuthRequest true "Данные для входа"
// @Success      200 {object} AuthResponse
// @Failure      400 {object} string "Неверный запрос"
// @Failure      500 {object} string "Внутренняя ошибка сервера"
// @Router       /auth [post]
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
		Message: resp.Message,
		Token:   resp.Token,
		UserID:  resp.UserId,
	})
}
