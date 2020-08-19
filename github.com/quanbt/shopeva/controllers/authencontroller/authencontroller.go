package authencontroller

import (
	"net/http"
	"shopeva/models"
	"shopeva/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//UserInputRequest ...
type UserInputRequest struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//Login ...
func Login(c *gin.Context) {
	var user models.User
	var userInput UserInputRequest
	var jwtService services.JWTService

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("username = ?", userInput.UserName).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tai khoan khong ton tai!"})
		return
	}

	match := CheckPasswordHash(userInput.Password, user.Pass)

	if !match {
		c.JSON(http.StatusOK, gin.H{"data": "Tai khoan hoac mat khau khong dung !"})
		return
	}

	token, err := jwtService.CreateAccessToken(user)
	tokens := map[string]string{
		"access_token":  token,
		"refresh_token": "balangnhang_werkj3@%#$&%^*dfwjerlkwje#$234123",
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": "token Errors !"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tokens})
}

//PassowrdHash ...
func PassowrdHash(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashed), err
}

//CheckPasswordHash ...
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false
	}

	return true
}
