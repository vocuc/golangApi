package productcontroler

import (
	"net/http"

	"shopeva/models"

	"github.com/gin-gonic/gin"
)

//Products Get all Products
func Products(c *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}
