package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"store/internal/dto"
	"store/internal/service/client"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ClientHandler struct {
	clientService service.ClientService
}

func NewClientHandler(clientService service.ClientService) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
	}
}

// @Summary      Создание клиента
// @Description  Добавляет нового клиента в систему. Формат поля "birthday": "1994-01-02T00:00:00Z",
// @Tags	     Клиенты
// @Param        request body dto.CreateClientRequest true "Данные клиента"
// @Success      201 {object} domain.Client "Клиент создан"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /clients/ [post]
func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateClientRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	client, err := h.clientService.CreateClient(r.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidClientData):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusCreated, client)
}

// @Summary      Удаление клиента
// @Description  Удаление клиента по ID
// @Tags	     Клиенты
// @Param        id path string true "UUID клиента"
// @Success      204 "Клиент удален"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /clients/{id} [delete]
func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	clientID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.clientService.DeleteClient(r.Context(), clientID)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrClientNotFound):
			writeError(w, http.StatusNotFound, err.Error())

		case errors.Is(err, service.ErrInvalidID):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusNoContent, fmt.Sprintf("client with id %s deleted", idStr))

}

// @Summary      Получить список клиентов по имени и фамилии
// @Description  Имя и фамилия обязательны
// @Tags	     Клиенты
// @Param        client_name query string true "Имя"
// @Param        client_surname query string true "Фамилия"
// @Success      200 "Данные успешно получены"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Router       /clients/by_name [get]
func (h *ClientHandler) GetClientsByFullName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("client_name")
	surname := r.URL.Query().Get("client_surname")

	clients, err := h.clientService.GetClientsByFullName(r.Context(), &dto.GetClientsByNameRequest{Name: name, SurName: surname})
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, clients)
}

// @Summary      Получить список клиентов с пагинацией
// @Description  При отсутствии поля Limit будет выдан весь список целиком
// @Tags	     Клиенты
// @Param        limit query string false "Лимит записей"
// @Param        offset query string false "Смещение"
// @Success      200 "Данные успешно получены"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /clients/by_page [get]
func (h *ClientHandler) GetClientsWithPage(w http.ResponseWriter, r *http.Request) {
	var limit *int32 = nil
	var offset *int32 = nil
	l, o := 0, 0
	var err error

	lStr := r.URL.Query().Get("limit")
	oStr := r.URL.Query().Get("offset")

	if lStr != "" {
		l, err = strconv.Atoi(lStr)
		l32 := int32(l)
		limit = &l32
	}

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if oStr != "" {
		o, err = strconv.Atoi(oStr)
		o32 := int32(o)
		offset = &o32
	}

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	clients, err := h.clientService.GetClientsWithPagination(r.Context(), &dto.GetclientsWithPaginationRequest{Limit: limit, Offset: offset})
	if err != nil {
		if errors.Is(err, service.ErrInvalidPagination) {
			writeError(w, http.StatusBadRequest, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusOK, clients)

}

// @Summary      Обновить адрес клиента
// @Tags	     Клиенты
// @Param		 id path string true "UUID клиента"
// @Param        request body dto.UpdateAddressParamsRequest true "Новый адрес"
// @Success      200 "Данные успешно обновлены"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /clients/{id}/address [patch]
func (h *ClientHandler) UpdateClientAddr(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	clientID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid client id")
		return
	}

	var req dto.UpdateAddressParamsRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	client, err := h.clientService.UpdateClientAddr(r.Context(), clientID, &req)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrClientNotFound):
			writeError(w, http.StatusNotFound, err.Error())

		case errors.Is(err, service.ErrInvalidAddrData):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusOK, client)
}
