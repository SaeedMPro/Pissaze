package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pissaze/internal/dto"
	"github.com/pissaze/internal/middleware"
	"github.com/pissaze/internal/service"
)

// /api/product/list
// /api/product/compatible
func registerProductRoutes(r *gin.Engine) {
	group := r.Group("/api/product")

	// Public endpoints
	group.GET("/list", getList)

	// VIP-only endpoints
	vipGroup := group.Group("")
	vipGroup.Use(middleware.Auth())
	{
		vipGroup.POST("/compatible", getCompatibleWithProductsList)
	}
}

// getList godoc
// @Summary List products
// @Description Get filtered list of products
// @Tags products
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=dto.ProductList}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/product/list [get]
func getList(c *gin.Context) {
	//TODO: filtering logic ???

	list, err := service.GETAllProducts()
	if err != nil {
		c.JSON(http.StatusOK, dto.ErrorResponse{
			Success: false,
			Error:   fmt.Sprintf("Error in fetching product's -> %s ", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    list,
	})
}

// getCompatibleWithProductsList godoc
// @Summary Find compatible products
// @Description Get products compatible with specified items
// @Tags products
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body dto.CompatibleRequest true "Product IDs and filters"
// @Success 200 {object} dto.SuccessResponse{data=[]models.Product}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/product/compatible [post]
func getCompatibleWithProductsList(c *gin.Context) {

	isVIP, exists := c.Get("is_vip")
	if !exists || !isVIP.(bool) {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{
			Success: false,
			Error:   "VIP access required",
		})
		return
	}

	var req dto.CompatibleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   fmt.Sprintf("Invalid request format: %v", err),
		})
		return
	}

	if len(req.ProductsID) == 0 {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Error:   "At least one product ID is required",
		})
		return
	}

	compatible, err := service.FindCompatibleWithCarts(req.ProductsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   fmt.Sprintf("Compatibility check failed: %v", err),
		})
		return
	}
	compatible = service.FilterBy(compatible,req.Filter)
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "Compatible products retrieved successfully",
		Data:    compatible,
	})
}
