package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pissaze/internal/dto"
	"github.com/pissaze/internal/service"
)

// /api/product/list
// /api/product/compatible
func registerProductRoutes(r *gin.Engine) {
	group := r.Group("/api/product")

	group.GET("/list", getList)
	group.POST("/compatible", getCompatibleWithProductsList)
}

func getList(c *gin.Context) {
	//TODO: filtering logic ???
	
	list, err := service.GETAllProducts()
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResponse{
			Success: false,
			Error: fmt.Sprintf("Error in fecthing product's -> %s ", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    list,
	})
}

func getCompatibleWithProductsList(c *gin.Context) {
	var req dto.CompatibleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "Invalid request format",
		})
		return
	}
	competible, err := service.FindcompatibleWithCarts(req.ProuductsID)
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResponse{
			Success: false,
			Error: fmt.Sprintf("Error in compatible product's -> %s ", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "get compatible product with your cart",
		Data:    competible,
	})
}