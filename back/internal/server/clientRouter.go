package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pissaze/internal/dto"
	"github.com/pissaze/internal/service"
)

// /
// /api/login
// /api/client/  --black
// /api/client/discountCode -- blue
// /api/client/cart  -- red

func registerClientRoutes(r *gin.Engine) {
	group := r.Group("/api/client")

	group.GET("/", getInfo)
	group.GET("/discountCode", getDiscounts)
	group.GET("/cart", getCart)
}

// api/client/
func getInfo(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error:   "Invalid request format",
		})
		return
	}

	client, err := service.GetClientByPhoneNumber(req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    client,
	})
}

func getDiscounts(c *gin.Context) {

}

func getCart(c *gin.Context) {

}
