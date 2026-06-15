package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"store/internal/dto"
	"store/internal/service/category"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// @Summary      Создание категории
// @Tags	     Категории
// @Security     BearerAuth
// @Param        request body dto.CreateCategoryRequest true "Название категории"
// @Success      201 {object} domain.Category "Категория создана"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /categories/ [post]
func (c *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	category, err := c.categoryService.CreateCategory(r.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidCategoryData):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusCreated, category)
}

// @Summary      Получить список всех категорий
// @Tags	     Категории
// @Security     BearerAuth
// @Success      200 "Данные успешно получены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /categories/ [get]
func (c *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.categoryService.GetAllCategories(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, categories)
}
