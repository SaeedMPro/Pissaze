package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pissaze/internal/dto"
	"github.com/pissaze/internal/middleware"
	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/service"
	"github.com/pissaze/internal/util"
)

// /
// /api/login
// /api/client/
// /api/client/discountCode
// /api/client/cart

func registerClientRoutes(r *gin.Engine) {
	group := r.Group("/api/client")
	group.Use(middleware.Auth())
	group.GET("/", getInfo)
	group.GET("/discountCode", getDiscounts)
	group.GET("/cart", getCart)
	group.GET("/lockcart", getLockCart)
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
	
	req, exist := c.Get("phone_number")
	reqString, ok := req.(string)
	if !exist || !ok{
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   "Key dosn't set correctly",
		})
		return
	}

	client, err := service.GetClientByPhoneNumber(reqString)
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

func getCart(c *gin.Context) {
	client, err := retriveUserByPhone(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	

	carts, err := service.GetClientCart(client.GetClient().ClientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	

	carts = util.NilFixer(carts)
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    carts,
	})
}

func getLockCart(c *gin.Context) {
	client, err := retriveUserByPhone(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	carts ,err :=service.GetClientSummuryOfCarts(client.GetClient().ClientID,5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	carts = util.NilFixer(carts)
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    carts,
	})
}

func getDiscounts(c *gin.Context) {
	
}


//------------------------- helper ----------------------------
func retriveUserByPhone(c *gin.Context)(models.ClientAbstract, error){
	req, exist := c.Get("phone_number")
	reqString, ok := req.(string)
	if !exist || !ok{
		return nil, errors.New("Key dosn't set correctly")
	}

	client, err := service.GetClientByPhoneNumber(reqString)
	if err != nil {
		return nil, err
	}

	return client, nil
}
