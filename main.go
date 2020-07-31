package main

import (
	"controllers/logincontroller"
	"controllers/productcontroller"
	"middlewares/authentoken"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	group := router.Group("/api/v1", authentoken.TokenAuthMiddleware())
	group.GET("/products", productcontroller.Show)
	group.POST("/login", logincontroller.Login)
	router.Run()
}
