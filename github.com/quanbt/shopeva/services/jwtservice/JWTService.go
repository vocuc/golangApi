package jwtservice

import (
	"fmt"
	"net/http"
	"os"
	"shopeva/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

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

//TokenValid ...
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.Claims)

	if !ok && !token.Valid {
		return err
	}

	fmt.Println(claims)

	return nil
}

//VerifyToken ...
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

//ExtractToken ...
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}
