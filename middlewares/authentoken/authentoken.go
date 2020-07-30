package authentoken

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//TokenAuthMiddleware ...
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.FormValue("api_token")

		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "API token required"})
			return
		}

		if token != requiredToken {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid API token"})
			return
		}

		c.Next()
	}
}
