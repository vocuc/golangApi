package couponcontroller

import (
	"net/http"
	"shopeva/helpers"
	"shopeva/models"
	"strings"
	"time"

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

	couponName := helpers.RandomString(8)

	coupon := models.Coupon{
		Name:        couponName,
		Value:       50000,
		Limit:       1,
		ExpiredDate: time.Now().Unix() + (86400 * 30),
		CreatedBy:   1,
		ForUser:     1,
		ForPhone:    input.Phone,
		Status:      models.CouponStatusActive,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	transaction := models.DB.Begin()

	if err := transaction.Create(&coupon).Error; err != nil {
		transaction.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error_code": 1, "data": ""})
		return
	}

	couponUser := models.CouponUser{
		CouponID:  int(coupon.ID),
		FullName:  input.Name,
		Phone:     input.Phone,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := transaction.Create(&couponUser).Error; err != nil {
		transaction.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error_code": 1, "data": ""})
		return
	}

	transaction.Commit()

	data := map[string]string{
		"coupon_name": strings.ToUpper(couponName),
		"coupon_id":   string(coupon.ID),
	}

	c.JSON(http.StatusBadRequest, gin.H{"error_code": 0, "data": data})
}
