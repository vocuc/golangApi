package services

import (
	"os"
	"shopeva/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTInterface ...
type JWTInterface interface {
	CreateAccessToken(user models.User) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

//JWTService ...
type JWTService struct {
}

//CreateAccessToken ...
func (s JWTService) CreateAccessToken(user models.User) (string, error) {
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
