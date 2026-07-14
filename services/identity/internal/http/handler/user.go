package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/dto"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/http/middleware"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/service"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/validator"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Me(c *gin.Context) {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing user identity"})
		return
	}

	user, err := h.svc.Get(c.Request.Context(), userID)
	if err != nil {
		writeUserError(c, err)
		return
	}

	c.JSON(http.StatusOK, toUserResponse(user))
}

func (h *UserHandler) UpdateMe(c *gin.Context) {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing user identity"})
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := validator.ValidateUpdateUser(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Name == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nothing to update"})
		return
	}

	user, err := h.svc.UpdateProfile(c.Request.Context(), userID, *req.Name)
	if err != nil {
		writeUserError(c, err)
		return
	}

	c.JSON(http.StatusOK, toUserResponse(user))
}

func toUserResponse(u *domain.User) dto.GetUserResponse {
	return dto.GetUserResponse{
		ID:              u.ID,
		Name:            u.Name,
		Email:           u.Email,
		EmailVerifiedAt: u.EmailVerifiedAt,
		CreatedAt:       u.CreatedAt,
	}
}

func writeUserError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
