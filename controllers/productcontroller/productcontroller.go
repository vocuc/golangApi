package productcontroller

import (
	"models"

	"github.com/gin-gonic/gin"
)

//Show ...
func Show(c *gin.Context) {
	var productModel models.ProductModel
	products, _ := productModel.FindAll()
	c.JSON(200, products)
}
