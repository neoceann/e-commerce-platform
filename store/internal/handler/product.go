package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"store/internal/dto"
	"store/internal/service/product"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// @Summary      Создание товара
// @Tags	     Товары
// @Security     BearerAuth
// @Param        request body dto.CreateProductRequest true "Данные товара"
// @Success      201 {object} domain.Product "Товар создан"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /products/ [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	product, err := h.productService.CreateProduct(r.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidProductData):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusCreated, product)
}

// @Summary      Добавление количества товара
// @Tags	     Товары
// @Security     BearerAuth
// @Param		 id path string true "UUID товара"
// @Param        request body dto.IncreaseProductStockRequest true "Количество"
// @Success      200 {object} domain.Product "Добавлено успешно"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /products/{id}/increase [patch]
func (h *ProductHandler) IncreaseProductStock(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	productID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var req dto.IncreaseProductStockRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
	}

	product, err := h.productService.IncreaseProductStock(r.Context(), productID, &dto.IncreaseProductStockRequest{Increasevalue: req.Increasevalue})

	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidID):
			writeError(w, http.StatusBadRequest, err.Error())

		case errors.Is(err, service.ErrDecreaseFailed):
			writeError(w, http.StatusBadRequest, err.Error())

		case errors.Is(err, service.ErrProductNotFound):
			writeError(w, http.StatusNotFound, err.Error())

		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusOK, product)
}

// @Summary      Уменьшение количества товара
// @Tags	     Товары
// @Security     BearerAuth
// @Param		 id path string true "UUID товара"
// @Param        request body dto.DecreaseProductStockRequest true "Количество"
// @Success      200 {object} domain.Product "Уменьшено успешно"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /products/{id}/decrease [patch]
func (h *ProductHandler) DecreaseProductStock(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	productID, err := uuid.Parse(idStr)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidID):
			writeError(w, http.StatusBadRequest, err.Error())

		case errors.Is(err, service.ErrDecreaseFailed):
			writeError(w, http.StatusBadRequest, err.Error())

		case errors.Is(err, service.ErrProductNotFound):
			writeError(w, http.StatusNotFound, err.Error())

		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	var req dto.DecreaseProductStockRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
	}

	product, err := h.productService.DecreaseProductStock(r.Context(), productID, &dto.DecreaseProductStockRequest{Decreasevalue: req.Decreasevalue})

	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidID), errors.Is(err, service.ErrDecreaseFailed):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusOK, product)
}

// @Summary      Получить товар по ID
// @Tags	     Товары
// @Security     BearerAuth
// @Param		 id path string true "UUID товара"
// @Success      200 {object} domain.Product "Данные получены успешно"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	productID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.productService.GetProductByID(r.Context(), productID)

	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusOK, product)
}

// @Summary      Удаление товара по ID
// @Tags	     Товары
// @Security     BearerAuth
// @Param        id path string true "UUID товара"
// @Success      204 "Товар удален"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /products/{id} [delete]
func (h *ProductHandler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	productID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.productService.DeleteProductByID(r.Context(), productID)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrProductNotFound):
			writeError(w, http.StatusNotFound, err.Error())

		case errors.Is(err, service.ErrInvalidID):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusNoContent, fmt.Sprintf("product with id %s deleted", idStr))
}

// @Summary      Получить список доступных товаров
// @Tags	     Товары
// @Security     BearerAuth
// @Success      200 "Данные успешно получены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /products/available [get]
func (h *ProductHandler) GetAvailableProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAvailableProducts(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, products)
}
