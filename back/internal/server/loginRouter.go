package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pissaze/internal/dto"
	"github.com/pissaze/internal/service"
)

// /api/login
func registerLoginRouter(r *gin.Engine) {
	group := r.Group("/api/login")

	group.POST("/", login)
}

// login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.SuccessResponse{data=dto.LoginRespond} "Login successful"
// @Failure 400 {object} dto.ErrorResponse "Invalid request format"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/login [post]
func login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "Invalid request format",
		})
		return
	}

	client, err := service.GetClientByPhoneNumber(req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	token, err := service.GenerateTokenStr(&client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "Login successful",
		Data: dto.LoginRespond{
			IsVip: client.IsVIP(),
			Token: token,
		},
	})
}
