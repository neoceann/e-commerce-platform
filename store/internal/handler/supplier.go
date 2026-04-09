package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"store/internal/dto"
	"store/internal/service/supplier"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type SupplierHandler struct {
	supplierService service.SupplierService
}

func NewSupplierHandler(supplierService service.SupplierService) *SupplierHandler {
	return &SupplierHandler{
		supplierService: supplierService,
	}
}

// @Summary      Создание поставщика
// @Description  Добавляет нового поставщика в систему
// @Tags	     Поставщики
// @Param        request body dto.CreateSupplierRequest true "Данные поставщика"
// @Success      201 {object} domain.Supplier "Поставщик создан"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /suppliers/ [post]
func (h *SupplierHandler) CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateSupplierRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	supplier, err := h.supplierService.CreateSupplier(r.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidSupplierData):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: " + err.Error())
		}
		return
	}
	
	writeJSON(w, http.StatusCreated, supplier)
}

// @Summary      Удаление поставщика
// @Description  Удаление поставщика по ID
// @Tags	     Поставщики
// @Param        id path string true "UUID поставщика"
// @Success      204 "Поставщик удален"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /suppliers/{id} [delete]
func (h *SupplierHandler) DeleteSupplier(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	supplierID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.supplierService.DeleteSupplier(r.Context(), supplierID)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrSupplierNotFound):
			writeError(w, http.StatusNotFound, err.Error())
		
		case errors.Is(err, service.ErrInvalidID):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: " + err.Error())
		}
		return
	}

	writeJSON(w, http.StatusNoContent, fmt.Sprintf("supplier with id %s deleted", idStr))
}

// @Summary      Получить список всех поставщиков
// @Tags	     Поставщики
// @Success      200 "Данные успешно получены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /suppliers/ [get]
func (h *SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := h.supplierService.GetAllSuppliers(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error: " + err.Error())
		return
	}
	
	writeJSON(w, http.StatusOK, suppliers)

}

// @Summary      Получить поставщика по ID
// @Tags	     Поставщики
// @Param        id path string true "UUID поставщика"
// @Success      200 "Данные успешно получены"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /suppliers/{id} [get]
func (h *SupplierHandler) GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	supplierID, err := uuid.Parse(idStr)

    if err != nil {
        writeError(w, http.StatusBadRequest, "invalid supplier id")
        return
    }

	supplier, err := h.supplierService.GetSupplierByID(r.Context(), supplierID)
	if err != nil {
		if errors.Is(err, service.ErrSupplierNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, "internal server error: " + err.Error())
		}
		return
	}
	
	writeJSON(w, http.StatusOK, supplier)

}

// @Summary      Обновить адрес поставщика
// @Tags	     Поставщики
// @Param		 id path string true "UUID поставщика"
// @Param        request body dto.UpdateAddressParamsRequest true "Новый адрес"
// @Success      200 "Данные успешно обновлены"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /suppliers/{id}/address [patch]
func (h *SupplierHandler) UpdateSupplierAddr(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	supplierID, err := uuid.Parse(idStr)

    if err != nil {
        writeError(w, http.StatusBadRequest, "invalid supplier id")
        return
    }

	var req dto.UpdateAddressParamsRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	supplier, err := h.supplierService.UpdateSupplierAddr(r.Context(), supplierID, &req)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrSupplierNotFound):
			writeError(w, http.StatusNotFound, err.Error())
		
		case errors.Is(err, service.ErrInvalidAddrData):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: " + err.Error())
		}
		return
	}

	writeJSON(w, http.StatusOK, supplier)
}