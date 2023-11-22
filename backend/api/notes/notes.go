package notes

import (
	"github.com/gin-gonic/gin"
	"github.com/meilisearch/meilisearch-go"
	"net/http"
	"os"
	"vortexnotes/backend/config"
	"vortexnotes/backend/database"
	"vortexnotes/backend/drivers"
)

func ListAllNotes(c *gin.Context) {
	var notes []database.Note
	database.DB.Select("id", "name", "content").Order("created_at desc").Limit(5).Find(&notes)
	c.JSON(http.StatusOK, notes)
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

func CreateNote(c *gin.Context) {
	type RequestData struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	var requestData RequestData

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := requestData.Name
	content := requestData.Content

	driver := drivers.LocalDriver{}
	err, note := driver.CreateNote(name, content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      note.ID,
		"name":    note.Name,
		"content": note.Content,
	})
}

func SearchNotes(c *gin.Context) {
	keywords := c.Query("keywords")
	notes, _ := config.MeiliSearchClient.Index("notes").Search(keywords, &meilisearch.SearchRequest{})

	type Result struct {
		Data     []interface{} `json:"data"`
		Duration float64       `json:"duration"`
	}

	result := Result{
		Data:     notes.Hits,
		Duration: float64(notes.ProcessingTimeMs) / 1000,
	}

	c.JSON(http.StatusOK, result)
}
