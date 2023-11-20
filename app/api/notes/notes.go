package notes

import (
	"github.com/gin-gonic/gin"
	"github.com/meilisearch/meilisearch-go"
	"net/http"
	"os"
	"vortexnotes/app/config"
	"vortexnotes/app/database"
)

func ListAllNotes(c *gin.Context) {
	notes, _ := config.MeiliSearchClient.Index("notes").Search("", &meilisearch.SearchRequest{
		Limit: 10,
	})

	c.JSON(http.StatusOK, notes.Hits)
}

func GetNote(c *gin.Context) {
	id := c.Param("id")

	var note database.Note
	var result = database.DB.First(&note, "id = ?", id)

	if result.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	fileData, err := os.ReadFile(config.LocalNotePath + note.Name)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	note.Content = string(fileData)

	c.JSON(http.StatusOK, note)
}

func SearchNotes(c *gin.Context) {
	keywords := c.Query("keywords")
	notes, _ := config.MeiliSearchClient.Index("notes").Search(keywords, &meilisearch.SearchRequest{})
	c.JSON(http.StatusOK, notes.Hits)
}
