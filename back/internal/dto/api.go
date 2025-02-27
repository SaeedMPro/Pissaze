package dto

import "github.com/pissaze/internal/models"

type ErrorResponse struct {
    Success bool   `json:"success" example:"false"`
    Error   string `json:"error" example:"error message"`
}

type SuccessResponse struct {
    Success bool        `json:"success" example:"true"`
    Message string      `json:"message,omitempty" example:"message"`
    Data    interface{} `json:"data,omitempty"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type ProductList struct {
	ProductList []models.ProductInterface
	Size        int
}
