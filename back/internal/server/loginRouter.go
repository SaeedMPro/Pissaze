package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pissaze/internal/dto"
	"github.com/pissaze/internal/service"
)

func registerLoginRouter(r *gin.Engine) {
	group := r.Group("/api/login")

	group.POST("/", login)
}

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

	token,err := service.GenerateTokenStr(&client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: false,
		Data: dto.LoginRespond{
			IsVip: client.IsVIP(),
			Token: token,
		},
	})
}