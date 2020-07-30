package main

import (
	"controllers/productcontroller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	group := router.Group("/api/v1")
	{
		group.GET("/products", productcontroller.Show)
	}

	router.Run()
}
