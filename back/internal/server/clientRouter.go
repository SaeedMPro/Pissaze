package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pissaze/internal/dto"
	"github.com/pissaze/internal/service"
)

// /
// /api/login
// /api/client/  
// /api/client/discountCode 
// /api/client/cart  

func registerClientRoutes(r *gin.Engine) {
	group := r.Group("/api/client")

	group.GET("/", getInfo)
	group.GET("/discountCode", getDiscounts)
	group.GET("/cart", getCart)
}

// getInfo godoc
// @Summary Get client information by phone number
// @Description Retrieve client details using their phone number. The phone number is provided in the request body. The response may include either a `Client` or a `VIPClient` object in the `data` field.
// @Tags client
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "phone_number"
// @Success 200 {object} dto.SuccessResponse{data=models.Client} "Client retrieved successfully"
// @Success 200 {object} dto.SuccessResponse{data=models.VIPClient} "VIP client retrieved successfully"
// @Failure 404 {object} dto.ErrorResponse "Client not found" 
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/client/ [GET]
func getInfo(c *gin.Context) {
	
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

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    client,
	})
}

func getDiscounts(c *gin.Context) {

}

func getCart(c *gin.Context) {

}
