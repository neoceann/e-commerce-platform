package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"store/internal/dto"
	"store/internal/service/image"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ImageHandler struct {
	imageService service.ImageService
}

func NewImageHandler(imageService service.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
	}
}

// @Summary      Создание изображения
// @Description  Добавляет новое изображение в систему
// @Tags	     Изображения товаров
// @Param        request body dto.CreateImageRequest true "Данные изображение"
// @Success      201 {object} domain.Image "Изображение создано"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /images/ [post]
func (h *ImageHandler) CreateImage(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateImageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body"+err.Error())
		return
	}

	image, err := h.imageService.CreateImage(r.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidImageData):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusCreated, dto.ImageToResponce(image))
}

// @Summary      Удаление изображение
// @Description  Удаление изображение по ID
// @Tags	     Изображения товаров
// @Param        id path string true "UUID изображения"
// @Success      204 "Изображение удалено"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /images/{id} [delete]
func (h *ImageHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	imageID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.imageService.DeleteImageByID(r.Context(), imageID)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrImageNotFound):
			writeError(w, http.StatusNotFound, err.Error())

		case errors.Is(err, service.ErrInvalidID):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusNoContent, fmt.Sprintf("image with id %s deleted", idStr))
}

// @Summary      Получить изображение по его ID
// @Tags	     Изображения товаров
// @Param        id path string true "UUID изображения"
// @Produce      application/octet-stream
// @Success      200 file binary "Изображение успешно получено"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /images/{id} [get]
func (h *ImageHandler) GetImageByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	imageID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid image id")
		return
	}

	image, err := h.imageService.GetImageByImageId(r.Context(), imageID)
	if err != nil {
		if errors.Is(err, service.ErrImageNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(image.ImageData)))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\"downloaded_image.png\"")
	w.Write(image.ImageData)
}

// @Summary      Получить список изображений по ID товара
// @Description  Для скачивания требуется запрашивать каждый ID отдельно (GET /images/{id})
// @Tags	     Изображения товаров
// @Param        id path string true "UUID товара"
// @Success      200 "Данные успешно получены"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure		 404 {object} map[string]string "Данные не найдены"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /images/product/{id} [get]
func (h *ImageHandler) GetImagesByProductID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	imageID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	images, err := h.imageService.GetImagesByProductID(r.Context(), imageID)
	if err != nil {
		if errors.Is(err, service.ErrImageNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
		} else {
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	resp := make([]*dto.ImageResponse, 0, len(images))

	for _, img := range images {
		resp = append(resp, dto.ImageToResponce(img))
	}

	writeJSON(w, http.StatusOK, resp)

}

// @Summary      Обновить изображение
// @Tags	     Изображения товаров
// @Param		 id path string true "UUID изображения"
// @Param        request body dto.UpdateImageRequest true "Новое изображение (base64)"
// @Success      200 "Данные успешно обновлены"
// @Failure      400 {object} map[string]string "Неверный запрос"
// @Failure      500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router       /images/{id}/update [patch]
func (h *ImageHandler) UpdateImage(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	imageID, err := uuid.Parse(idStr)

	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid image id")
		return
	}

	var req dto.UpdateImageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	image, err := h.imageService.UpdateImage(r.Context(), imageID, &req)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrImageNotFound):
			writeError(w, http.StatusNotFound, err.Error())

		case errors.Is(err, service.ErrInvalidImageData):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error: "+err.Error())
		}
		return
	}

	writeJSON(w, http.StatusOK, dto.ImageToResponce(image))
}
