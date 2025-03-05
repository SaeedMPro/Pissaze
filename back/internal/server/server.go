package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/pissaze/docs"
	"github.com/pissaze/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	host = "localhost"
	port = "8082"
)

func Start() {
	r := gin.Default()
	
	r.Use(middleware.CORS())
	registerRoutes(r)
	r.Run(fmt.Sprintf("%s:%s", host, port))
}

func registerRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	registerLoginRouter(r)
	registerClientRoutes(r)
	registerProductRoutes(r)
}
