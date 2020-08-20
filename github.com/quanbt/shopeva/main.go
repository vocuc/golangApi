package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"shopeva/models"
	"shopeva/routers"
)

var router *gin.Engine

func init() {
	os.Setenv("PORT", "9000")
	os.Setenv("GIN_MODE", "release")
	os.Setenv("ACCESS_SECRET", "jd4f3ls@jflks#$25fgsdf")
}

func main() {
	models.ConnectDataBase()
	routers.InitRoutes()
	defer models.DB.Close()
}
