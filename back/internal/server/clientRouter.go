package server

import "github.com/gin-gonic/gin"

// /
// /api/login
// /api/client/  --black
// /api/client/discountcode -- blue
// /api/client/cart  -- red

func registerClientRoutes(r *gin.Engine){

	group := r.Group("api/client")

	group.GET("/",getInfo)
	group.GET("/discountcode",getDiscounds)
	group.GET("/cart",getCart)
}

func getInfo(c *gin.Context){

}

func getDiscounds(c *gin.Context){

}

func getCart(c *gin.Context){

}

