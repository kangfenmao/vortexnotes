package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"vortexnotes/app/api/notes"
	"vortexnotes/app/api/website"
	"vortexnotes/app/config"
	"vortexnotes/app/database"
	"vortexnotes/app/web"
)

func Start() {
	database.InitializeDatabase()

	server := gin.Default()
	server.Use(cors.Default())

	server.GET("/", website.ServeRoot)
	server.GET("/assets/*filepath", website.ServeAssets)
	server.Static("/notes/attachments", config.LocalNotePath+"attachments")
	server.StaticFS("/public", http.FS(web.Favicon))
	server.NoRoute(website.NoRoot)

	api := server.Group("/api")
	{
		api.GET("/notes", notes.ListAllNotes)
		api.GET("/notes/:id", notes.GetNote)
		api.POST("/notes/new", notes.CreateNote)
		api.GET("/search", notes.SearchNotes)
	}

	server.SetTrustedProxies(nil)
	server.Run("127.0.0.1:7701")
}
