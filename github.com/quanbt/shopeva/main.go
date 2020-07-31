package main

import (
	"github.com/gin-gonic/gin"

	"shopeva/controllers/productcontroler"
	"shopeva/models"
)

func main() {
	router := gin.Default()

	models.ConnectDataBase()

	group := router.Group("/api/v1")

	group.GET("/products", productcontroler.Products)

	router.Run()
}
