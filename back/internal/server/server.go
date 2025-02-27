package server

import (
	"fmt"
	"net/http"


	"github.com/gin-gonic/gin"
	_ "github.com/pissaze/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	host = "localhost"
	port = "8082"
)


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
func Start() {
	r := gin.Default()
	
	r.Use(CORSMiddleware())
	registerRoutes(r)
	r.Run(fmt.Sprintf("%s:%s", host, port))
}

func registerRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	registerClientRoutes(r)
	registerProductRoutes(r)
}
