package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func HasPermission(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedPasscode := os.Getenv("VORTEXNOTES_PASSCODE")
		authorizationHeader := c.GetHeader("Authorization")
		passcode := strings.TrimPrefix(authorizationHeader, "Bearer ")

		needAuthScopes := os.Getenv("VORTEXNOTES_AUTH_SCOPE")
		if needAuthScopes == "" {
			needAuthScopes = "show,create,edit,delete"
		}

		if expectedPasscode == "" {
			c.Next()
			return
		}

		if strings.Contains(needAuthScopes, scope) {
			if passcode != expectedPasscode {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				return
			}
		}

		c.Next()
	}
}
