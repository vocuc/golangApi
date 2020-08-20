package couponcontroller

import (
	"github.com/gin-gonic/gin"
)

type couponInput struct {
	Phone string `json:"phone" binding:"required"`
}

//CreateaCoupon ...
func CreateaCoupon(c *gin.Context) {

}
