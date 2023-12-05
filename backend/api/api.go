package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"vortexnotes/backend/api/auth"
	"vortexnotes/backend/api/configuration"
	"vortexnotes/backend/api/middlewares"
	"vortexnotes/backend/api/notes"
	"vortexnotes/backend/api/website"
	"vortexnotes/backend/config"
)

func Start() {
	server := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	server.Use(cors.New(corsConfig))

	server.GET("/", website.ServeRoot)
	server.GET("/assets/*filepath", website.ServeAssets)
	server.GET("/public/*filepath", website.ServePublic)
	server.Static("/notes/attachments", config.LocalNotePath+"attachments")
	server.NoRoute(website.NoRoot)

	api := server.Group("/api")
	{
		api.GET("/config", configuration.Config)
		api.POST("/auth", auth.Auth)
		api.GET("/search", middlewares.HasPermission("show"), notes.SearchNotes)
		api.GET("/notes", middlewares.HasPermission("show"), notes.ListAllNotes)
		api.GET("/notes/:id", middlewares.HasPermission("show"), notes.GetNote)
		api.DELETE("/notes/:id", middlewares.HasPermission("delete"), notes.DeleteNote)
		api.PATCH("/notes/:id", middlewares.HasPermission("edit"), notes.UpdateNote)
		api.POST("/notes/new", middlewares.HasPermission("create"), notes.CreateNote)
	}

	server.SetTrustedProxies(nil)
	server.Run(config.ApiHost + ":" + config.ApiPort)
}
