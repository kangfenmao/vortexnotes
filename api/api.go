package api

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"vortexnotes/api/indexer"
	"vortexnotes/api/notes"
	"vortexnotes/config"
)

func Start() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile(config.WebRoot, true)))

	api := r.Group("/api")
	{
		api.GET("/api/notes", notes.ListAllNotes)
		api.GET("/api/search", notes.SearchNotes)
		api.POST("/api/indexes", indexer.StartIndex)
	}

	r.Run("0.0.0.0:6480")
}
