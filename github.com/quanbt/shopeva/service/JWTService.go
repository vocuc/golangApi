package service

import (
	"shopeva/models"

	"github.com/dgrijalva/jwt-go"
)

//JWTService ...
type JWTService interface {
	CreateAccessToken(user models.User) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

