package middleware

import (
	"net/http"
	"shopeva/services/jwtservice"

	"github.com/gin-gonic/gin"
)

//AuthMiddleware ...
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwtservice.TokenValid(c.Request)

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}
