package couponcontroller

import (
	"net/http"
	"shopeva/helpers"

	"github.com/gin-gonic/gin"
)

type couponInput struct {
	Phone string `json:"phone" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

//CreateaCoupon ...
func CreateaCoupon(c *gin.Context) {
	var input couponInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	c.JSON(http.StatusBadRequest, gin.H{"data": helpers.RandomString(10)})
}
