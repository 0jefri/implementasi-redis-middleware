package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/api/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Login user
// @Description Authenticate a user with their username and passwor.
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginRequest body LoginRequest true "Login request data"
// @Success 200 {object} httputil.HTTPError
// @Failure 400 {object} httputil.HTTPError
// @Failure 401 {object} httputil.HTTPError
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {

	input := LoginRequest{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}
