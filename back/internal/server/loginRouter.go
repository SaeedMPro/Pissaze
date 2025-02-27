package server

import "github.com/gin-gonic/gin"

func registerLoginRouter(r *gin.Engine) {
	group := r.Group("/api/login")

	group.POST("/", login)
}

func login(c *gin.Context) {
	//TODO
}