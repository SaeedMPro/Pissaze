package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pissaze/internal/dto"
	"github.com/pissaze/internal/models"
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

	var list []models.ProductInterface

	cpus, err := service.GetAllCPU()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	for _, tmp := range cpus {
		list = append(list, tmp)
	}

	hhds, err := service.GetAllHDD()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	for _, tmp := range hhds {
		list = append(list, tmp)
	}

	coolers, err := service.GetAllCooler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	for _, tmp := range coolers {
		list = append(list, tmp)
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    list,
	})
}

func getCompatibleWithProductsList(c *gin.Context) {
	//TODO
}