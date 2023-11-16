package notes

import (
	"github.com/gin-gonic/gin"
	"github.com/meilisearch/meilisearch-go"
	"net/http"
	"vortexnotes/config"
)

func ListAllNotes(c *gin.Context) {
	notes, _ := config.MeiliSearchClient.Index("notes").Search("", &meilisearch.SearchRequest{
		Limit: 10,
	})

	c.JSON(http.StatusOK, notes.Hits)
}

func SearchNotes(c *gin.Context) {
	keywords := c.Query("keywords")
	notes, _ := config.MeiliSearchClient.Index("notes").Search(keywords, &meilisearch.SearchRequest{})
	c.JSON(http.StatusOK, notes.Hits)
}
