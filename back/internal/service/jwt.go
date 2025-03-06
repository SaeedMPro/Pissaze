package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pissaze/internal/models"
)

var JwtSecretKey = []byte("your_secret_key")
const(
	expireTime   = 24 * time.Hour
) 


func GenerateTokenStr(client *models.ClientAbstract) (string, error) {
	exp := time.Now().Add(expireTime)
	claims := models.Claims{
		PhoneNumber: (*client).GetClient().PhoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt : &jwt.Time{Time: exp}, 
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractPhoneNumber(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims.PhoneNumber, nil
	}

	return "", errors.New("invalid token")
}



