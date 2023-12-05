package configuration

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Config(c *gin.Context) {
	needAuthScopes := os.Getenv("VORTEXNOTES_AUTH_SCOPE")
	passcode := os.Getenv("VORTEXNOTES_PASSCODE")
	authType := "none"

	if needAuthScopes == "" {
		needAuthScopes = "show,create,edit,delete"
	}

	if passcode != "" {
		authType = "passcode"
		c.JSON(http.StatusOK, gin.H{
			"auth_type":  authType,
			"auth_scope": needAuthScopes,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"auth_type": authType,
	})
}
