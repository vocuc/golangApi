package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"shopeva/controllers/productcontroler"
	"shopeva/models"
)

func main() {
	os.Setenv("PORT", "9000")

	router := gin.Default()

	models.ConnectDataBase()

	group := router.Group("/api/v1")
	group.GET("/products", productcontroler.Products)
	group.POST("/products", productcontroler.Store)
	group.GET("/products/:id", productcontroler.FindProduct)
	group.PUT("/products/:id", productcontroler.Update)
	router.Run()
}
