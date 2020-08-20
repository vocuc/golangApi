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

//ShowCoupon ...
func ShowCoupon(c *gin.Context) {
	var coupon models.Coupon

	if err := models.DB.Where("id = ?", c.Param("id")).First(&coupon).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": coupon})
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
		Name:        strings.ToUpper(couponName),
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

	c.JSON(http.StatusBadRequest, gin.H{"error_code": 0, "data": coupon})
}
