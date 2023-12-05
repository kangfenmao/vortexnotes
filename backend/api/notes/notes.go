package notes

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"vortexnotes/backend/blevesearch"
	"vortexnotes/backend/config"
	"vortexnotes/backend/database"
	"vortexnotes/backend/indexer"
)

func ListAllNotes(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	Order := c.DefaultQuery("sort", "name asc")
	Order = strings.Replace(Order, ":", " ", 1)
	offset := (page - 1) * limit

	var notes []database.Note
	database.DB.Order(Order).Limit(limit).Offset(offset).Find(&notes)

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
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	name := requestData.Name
	content := requestData.Content

	note, err := indexer.CreateNote(name, content)
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
	limitQuery := c.Query("limit")

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		limit = 50
	}

	query := bleve.NewQueryStringQuery(keywords)
	searchRequest := bleve.NewSearchRequestOptions(query, limit, 0, false)
	searchRequest.Highlight = bleve.NewHighlight()
	searchResult, err := blevesearch.NotesIndex.Search(searchRequest)
	if err != nil {
		fmt.Println("Search failed:", err)
		return
	}

	var notes []database.Note
	for _, hit := range searchResult.Hits {
		id := hit.ID
		var note database.Note
		database.DB.First(&note, "id = ?", id)

		if hit.Fragments["name"] != nil {
			note.Name = hit.Fragments["name"][0]
		}

		if hit.Fragments["content"] != nil {
			note.Content = hit.Fragments["content"][0]
		}

		notes = append(notes, note)
	}

	type Result struct {
		Data      []database.Note `json:"data"`
		Duration  float64         `json:"duration"`
		Page      int64           `json:"page"`
		TotalPage uint64          `json:"total_page"`
	}

	if notes == nil {
		notes = []database.Note{}
	}

	c.JSON(http.StatusOK, Result{
		Data:      notes,
		Duration:  searchResult.Took.Seconds(),
		TotalPage: searchResult.Total,
	})
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")

	err := indexer.DeleteNote(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

func UpdateNote(c *gin.Context) {
	id := c.Param("id")

	err := indexer.DeleteNote(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}

	CreateNote(c)
}
