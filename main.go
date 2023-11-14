package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vortex-notes/indexer"
)

func main() {
	indexer.Start()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:6480")
}
