package authencontroller

import (
	"net/http"
	"os"
	"shopeva/models"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	token, err := CreateAccessToken(user)
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

//CreateAccessToken ...
func CreateAccessToken(user models.User) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = user.ID
	atClaims["userName"] = user.Name
	atClaims["store_id"] = user.StoreID
	atClaims["admin_level"] = user.AdminLevel
	atClaims["workGroup"] = user.WorkGroup
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	accesstoken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := accesstoken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
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
