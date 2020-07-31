package authentoken

import (
	"github.com/gin-gonic/gin"
)

//TokenAuthMiddleware ...
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.Request.FormValue("api_token")

		// if token == "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "API token required"})
		// 	c.Abort()
		// }

		c.Next()
	}
}
