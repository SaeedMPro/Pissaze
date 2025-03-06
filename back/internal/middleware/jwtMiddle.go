package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pissaze/internal/dto"
	"github.com/pissaze/internal/service"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Error: "token not found",
			})
			c.Abort()
			return
		}

		if len(token) > 6 && token[:7] == "Bearer " {
			token = token[7:] 
		}

		phoneNumber, err := service.ExtractPhoneNumber(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Error: fmt.Sprintf("can pars token : %s", err.Error()),
			})
			c.Abort()
			return
		}

		c.Set("phone_number", phoneNumber) 
		c.Next()
	}
}
