package main

import (
	"os"
	"shopeva/controllers/authencontroller"

	"github.com/gin-gonic/gin"

	"shopeva/controllers/productcontroler"
	"shopeva/models"
)

func init() {
	os.Setenv("PORT", "9000")
	os.Setenv("GIN_MODE", "release")
	os.Setenv("ACCESS_SECRET", "jd4f3ls@jflks#$25fgsdf")
}

func main() {
	router := gin.Default()
	models.ConnectDataBase()
	group := router.Group("/api/v1")

	//Product
	group.GET("/products", productcontroler.Products)
	group.POST("/products", productcontroler.Store)
	group.GET("/products/:id", productcontroler.FindProduct)
	group.PUT("/products/:id", productcontroler.Update)

	//Login
	group.POST("/login", authencontroller.Login)

	router.Run()

	defer models.DB.Close()
}
