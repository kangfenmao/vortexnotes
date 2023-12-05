package configuration

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Config(c *gin.Context) {
	needAuthScopes := os.Getenv("VORTEXNOTES_AUTH_SCOPE")
	passcode := os.Getenv("VORTEXNOTES_PASSCODE")
	auth := "none"

	if needAuthScopes == "" {
		needAuthScopes = "show,create,edit,delete"
	}

	if passcode != "" {
		auth = "passcode"
	}

	c.JSON(http.StatusOK, gin.H{
		"auth_scope": needAuthScopes,
		"auth":       auth,
	})
}
