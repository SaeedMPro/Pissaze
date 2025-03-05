package models

import "github.com/dgrijalva/jwt-go/v4"

type Claims struct {
	PhoneNumber string  `json:"phone_number" binding:"required"`
	jwt.StandardClaims 
}