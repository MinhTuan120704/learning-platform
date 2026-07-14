package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/dto"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/service"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/validator"
)

type CategoryHandler struct {
	svc *service.CategoryService
}

func NewCategoryHandler(svc *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: svc}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := validator.ValidateCreateCategory(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.svc.Create(c.Request.Context(), req.Name, req.Slug, req.Description)
	if err != nil {
		writeCategoryError(c, err)
		return
	}

	c.JSON(http.StatusCreated, toCategoryResponse(category))
}

func (h *CategoryHandler) List(c *gin.Context) {
	categories, err := h.svc.List(c.Request.Context())
	if err != nil {
		writeCategoryError(c, err)
		return
	}

	res := make([]dto.CategoryResponse, 0, len(categories))
	for _, cat := range categories {
		res = append(res, toCategoryResponse(&cat))
	}
	c.JSON(http.StatusOK, res)
}

func (h *CategoryHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	category, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		writeCategoryError(c, err)
		return
	}

	c.JSON(http.StatusOK, toCategoryResponse(category))
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := validator.ValidateUpdateCategory(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.svc.Update(c.Request.Context(), id, req.Name, req.Slug, req.Description)
	if err != nil {
		writeCategoryError(c, err)
		return
	}

	c.JSON(http.StatusOK, toCategoryResponse(category))
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		writeCategoryError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func toCategoryResponse(c *domain.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:          c.ID,
		Name:        c.Name,
		Slug:        c.Slug,
		Description: c.Description,
	}
}

func writeCategoryError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrCategoryNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrCategorySlugAlreadyExists):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrCategoryNameRequired), errors.Is(err, domain.ErrCategorySlugRequired):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
