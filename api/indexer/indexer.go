package indexer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vortexnotes/indexer"
)

func StartIndex(c *gin.Context) {
	indexer.Start()

	c.JSON(http.StatusOK, gin.H{
		"message": "Starting indexer",
	})
}
