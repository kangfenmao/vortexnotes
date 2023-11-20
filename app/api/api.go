package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"vortexnotes/app/api/indexer"
	"vortexnotes/app/api/notes"
	"vortexnotes/app/config"
	"vortexnotes/app/database"
)

func Start() {
	database.InitializeDatabase()

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile(config.WebRoot, true)))
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/notes", notes.ListAllNotes)
		api.GET("/notes/:id", notes.GetNote)
		api.GET("/search", notes.SearchNotes)
		api.POST("/indexes", indexer.StartIndex)
	}

	r.Run("0.0.0.0:7701")
}
