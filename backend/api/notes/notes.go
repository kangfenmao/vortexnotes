package notes

import (
	"github.com/gin-gonic/gin"
	"github.com/meilisearch/meilisearch-go"
	"net/http"
	"os"
	"strconv"
	"vortexnotes/backend/config"
	"vortexnotes/backend/database"
	"vortexnotes/backend/indexer"
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

	err, note := indexer.LocalIndexer{}.CreateNote(name, content)
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
	pageQuery := c.Query("page")
	limitQuery := c.Query("limit")

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		limit = 20
	}

	offset := int64((page - 1) * limit)
	searchResults, _ := config.MeiliSearchClient.Index("notes").Search(keywords, &meilisearch.SearchRequest{
		Offset:                offset,
		Limit:                 int64(limit),
		AttributesToHighlight: []string{"content"},
	})

	type Result struct {
		Data      []interface{} `json:"data"`
		Duration  float64       `json:"duration"`
		Page      int64         `json:"page"`
		TotalPage int64         `json:"total_page"`
	}

	result := Result{
		Data:      searchResults.Hits,
		Duration:  float64(searchResults.ProcessingTimeMs) / 1000,
		Page:      offset,
		TotalPage: searchResults.EstimatedTotalHits,
	}

	c.JSON(http.StatusOK, result)
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")

	err := indexer.LocalIndexer{}.DeleteNote(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

func UpdateNote(c *gin.Context) {
	id := c.Param("id")

	err := indexer.LocalIndexer{}.DeleteNote(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	CreateNote(c)
}
