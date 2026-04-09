package handler

import (
	"errors"
	"net/http"
	"store/internal/service/address"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type AddressHandler struct {
	addressService service.AddressService
}

func NewAddressHandler(a service.AddressService) *AddressHandler {
	return &AddressHandler{
		addressService: a,
	}
}

// @Summary      Получить адрес по ID
// @Tags	     Адреса
// @Param		 id path string true "UUID адреса"
// @Success      200 {object} domain.Product "Данные получены успешно"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /addresses/{id} [get]
func (h *AddressHandler) GetAddressByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	addrID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.addressService.GetAddressByID(r.Context(), addrID)

	if err != nil {
		if errors.Is(err, service.ErrAddrNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, "internal server error: " + err.Error())
		}
		return
	}
	
	writeJSON(w, http.StatusOK, product)
}