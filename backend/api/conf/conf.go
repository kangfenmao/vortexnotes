package conf

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Config(c *gin.Context) {
	authScopes := os.Getenv("VORTEXNOTES_AUTH_SCOPE")
	passcode := os.Getenv("VORTEXNOTES_PASSCODE")

	if authScopes == "" {
		authScopes = "show,create,edit,delete"
	}

	if passcode != "" {
		c.JSON(http.StatusOK, gin.H{
			"auth_type":  "passcode",
			"auth_scope": authScopes,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"auth_type": "none",
	})
}
