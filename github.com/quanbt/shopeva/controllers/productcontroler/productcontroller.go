package productcontroler

import (
	"database/sql"
	"net/http"

	"shopeva/models"

	"github.com/gin-gonic/gin"
)

type ProductFormInput struct {
	ProductName   string `json:"name" binding:"required"`
	ProductPrice  int    `json:"price" binding:"required"`
	ProductAvarta string `json:"avarta" binding:"required"`
}

//Products Get all Products
func Products(c *gin.Context) {

	var products []models.Product

	models.DB.Limit(10).Offset(0).Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

//FindProduct Tìm sản phẩm theo ID
func FindProduct(c *gin.Context) {
	var product models.Product

	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

//Store Tạo sản phẩm mới
func Store(c *gin.Context) {
	var input ProductFormInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name: sql.NullString{
			String: input.ProductName,
			Valid:  true,
		},
		Price:  input.ProductPrice,
		Images: input.ProductAvarta,
	}

	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

//UpdateProductInput ...
type UpdateProductInput struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Images string `json:"images"`
}

//Update ...
func Update(c *gin.Context) {
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&product).Updates(&models.Product{
		Name:   sql.NullString{String: input.Name, Valid: true},
		Price:  input.Price,
		Images: input.Images,
	})

	c.JSON(http.StatusOK, gin.H{"data": product})
}
