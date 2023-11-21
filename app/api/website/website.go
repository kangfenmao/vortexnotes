package website

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vortexnotes/app/web"
)

func ServeRoot(c *gin.Context) {
	http.FileServer(http.FS(web.IndexHtml)).ServeHTTP(c.Writer, c.Request)
}

func ServeAssets(c *gin.Context) {
	http.FileServer(http.FS(web.Assets)).ServeHTTP(c.Writer, c.Request)
}

func NoRoot(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(web.IndexByte)
	c.Writer.Header().Add("Accept", "text/html")
	c.Writer.Flush()
}
