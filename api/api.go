package api

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"vortexnotes/api/indexer"
	"vortexnotes/api/notes"
	"vortexnotes/app/config"
)

func Start() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile(config.WebRoot, true)))

	api := r.Group("/api")
	{
		api.GET("/notes", notes.ListAllNotes)
		api.GET("/search", notes.SearchNotes)
		api.POST("/indexes", indexer.StartIndex)
	}

	r.Run("0.0.0.0:7701")
}
