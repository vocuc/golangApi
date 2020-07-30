package productcontroller

import (
	"models"

	"github.com/gin-gonic/gin"
)

//Show ...
func Show(c *gin.Context) {

	var productModel models.ProductModel
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	products, _ := productModel.FindAll(limit, offset)

	c.JSON(200, products)
}
