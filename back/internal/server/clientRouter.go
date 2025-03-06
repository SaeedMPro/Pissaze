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

// /api/client/
// /api/client/discountCode
// /api/client/cart
// /api/client/cart/lockCart
func registerClientRoutes(r *gin.Engine) {
	group := r.Group("/api/client")
	group.Use(middleware.Auth())

	group.GET("/", getInfo)
	group.GET("/discountCode", getDiscounts)
	group.GET("/cart", getCart)
	group.GET("/lockCart", getLockCart)
}

// getInfo godoc
// @Summary Get client information
// @Description Retrieve client details using JWT token
// @Tags client
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=models.Client} "Client retrieved successfully"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/client/ [get]
func getInfo(c *gin.Context) {
	
	req, exist := c.Get("phone_number")
	reqString, ok := req.(string)
	if !exist || !ok{
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   "Key doesn't set correctly",
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

// getLockCart godoc
// @Summary Get locked cart summary
// @Description Retrieve summary of locked carts within specified days
// @Tags client
// @Security ApiKeyAuth
// @Produce json
// @Param days query int false "Number of days to look back (default 5)"
// @Success 200 {object} dto.SuccessResponse{data=[]models.LockedShoppingCart} "Locked carts retrieved successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid days parameter"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/client/lockCart [get]
func getCart(c *gin.Context) {
	client, err := retrieveUserByPhone(c)
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
	client, err := retrieveUserByPhone(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	//TODO: make days query param
	carts ,err := service.GetClientSummaryOfCarts(client.GetClient().ClientID,5)
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

// getDiscounts godoc
// @Summary Get client's discount codes
// @Description Retrieve all active discount codes for the client
// @Tags client
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=dto.DiscountRespond} "Discount codes retrieved successfully"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /api/client/discountCode [get]
func getDiscounts(c *gin.Context) {
	client, err := retrieveUserByPhone(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	codes, err := service.GetClientPrivateCode(client,7)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	count, err := service.NumberOfGiftedCode(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	res := dto.DiscountRespond{
		NumberOfGiftCode: count,
		DiscountCodes: util.NilFixer(codes),
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    res,
	})
}


//------------------------- helper ----------------------------
func retrieveUserByPhone(c *gin.Context)(models.ClientAbstract, error){
	req, exist := c.Get("phone_number")
	reqString, ok := req.(string)
	if !exist || !ok{
		return nil, errors.New("key doesn't set correctly")
	}

	client, err := service.GetClientByPhoneNumber(reqString)
	if err != nil {
		return nil, err
	}

	return client, nil
}
