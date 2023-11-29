package website

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vortexnotes/backend/web"
)

func ServeRoot(c *gin.Context) {
	http.FileServer(http.FS(web.IndexHtml)).ServeHTTP(c.Writer, c.Request)
}

func ServeAssets(c *gin.Context) {
	c.Header("Cache-Control", "public, max-age=31536000")
	http.FileServer(http.FS(web.Assets)).ServeHTTP(c.Writer, c.Request)
}

func ServePublic(c *gin.Context) {
	c.Header("Cache-Control", "public, max-age=86400")
	http.FileServer(http.FS(web.Public)).ServeHTTP(c.Writer, c.Request)
}

func NoRoot(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(web.IndexByte)
	c.Writer.Header().Add("Accept", "text/html")
	c.Writer.Flush()
}
