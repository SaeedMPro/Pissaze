package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pissaze/internal/models"
)

var JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
const(
	expireTime   = 24 * time.Hour
) 

func GenerateTokenStr(client *models.ClientAbstract) (string, error) {

    fmt.Println(JwtSecretKey)
    exp := time.Now().Add(expireTime)
    claims := models.Claims{
        PhoneNumber: (*client).GetClient().PhoneNumber,
        IsVIP:       (*client).IsVIP(), 
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: &jwt.Time{Time: exp},
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    tokenString, err := token.SignedString(JwtSecretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ExtractClaims(tokenString string) (*models.Claims, error) {
    claims := &models.Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return JwtSecretKey, nil
    })
    
    if err != nil {
        return nil, err
    }
    
    if !token.Valid {
        return nil, errors.New("invalid token")
    }
    
    return claims, nil
}