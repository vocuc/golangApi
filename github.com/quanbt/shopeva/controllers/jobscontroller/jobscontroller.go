package jobscontroller

import (
	"net/http"
	"shopeva/models"
	"time"

	"github.com/gin-gonic/gin"
)

type jobInput struct {
	FullName     string `json:"name" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Cast         int    `json:"cast" binding:"required"`
	LinkVideo    string `json:"video" binding:"required"`
	LinkFacebook string `json:"facebook" binding:"required"`
}

//Store Tạo sản phẩm mới
func Store(c *gin.Context) {
	var input jobInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job := models.Job{
		FullName:     input.FullName,
		Phone:        input.Phone,
		Cast:         input.Cast,
		LinkVideo:    input.LinkVideo,
		LinkFacebook: input.LinkFacebook,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	result := models.DB.Create(&job)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": job})
}
