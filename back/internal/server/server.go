package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	host = "localhost"
	port = "8082"
)

func Start() {
	r := gin.Default()
	registerRoutes(r)
	r.Run(fmt.Sprintf("%s:%s", host, port))
}

func registerRoutes(r *gin.Engine) {
	registerClientRoutes(r)
	registerProductRoutes(r)
}
