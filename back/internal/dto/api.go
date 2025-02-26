package dto

import "github.com/pissaze/internal/models"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type ProductList struct {
	ProductList []models.ProductInterface
	Size        int
}
