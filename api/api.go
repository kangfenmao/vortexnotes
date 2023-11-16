package api

import (
	"github.com/gin-gonic/gin"
	"vortexnotes/api/indexer"
	"vortexnotes/api/notes"
)

func Start() {
	r := gin.Default()
	r.GET("/", notes.ListAllNotes)
	r.GET("/search", notes.SearchNotes)
	r.POST("/indexes", indexer.StartIndex)
	r.Run("0.0.0.0:6480")
}
