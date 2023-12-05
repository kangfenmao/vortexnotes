package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Auth(c *gin.Context) {
	passcode := os.Getenv("VORTEXNOTES_PASSCODE")

	type RequestData struct {
		Passcode string `json:"passcode"`
	}

	var requestData RequestData

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if passcode != "" {
		if requestData.Passcode != passcode {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Passcode Invalid"})
			return
		}
	}

	authScopes := os.Getenv("VORTEXNOTES_AUTH_SCOPE")
	if authScopes == "" {
		authScopes = "show,create,edit,delete"
	}

	c.JSON(http.StatusOK, gin.H{
		"auth_scope": authScopes,
	})
}
