package local

import (
	"github.com/goccy/go-json"
	"github.com/meilisearch/meilisearch-go"
	"io"
	"os"
	"vortex-notes/indexer/config"
	"vortex-notes/indexer/logger"
	"vortex-notes/indexer/sqlite"
)

type Driver struct {
}

func (local Driver) ListNotes() []string {
	err := CreateDirectoryIfNotExists(config.LocalNotePath)
	if err != nil {
		var list []string
		logger.Logger.Fatal("Error:", err)
		return list
	}

	var notes, _ = ListTextFiles(config.LocalNotePath)
	return notes
}

func (local Driver) AddNoteToDatabase(path string) {
	id, _ := CalculateFileHash(path)

	fileInfo, err := os.Stat(path)
	if err != nil {
		logger.Logger.Fatal("Stat File Error:", err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		logger.Logger.Fatal("ReadFile Error:", err)
	}

	err = sqlite.InsertNote(id, fileInfo.Name(), content)
	if err != nil {
		logger.Logger.Fatal("InsertNote Error:", err)
	}
}

func (local Driver) SyncNoteToMeiliSearch() {
	sqlite.GenerateNotesJsonFile()

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700",
		APIKey: "zXEpbeyeGtGi8DQbfOSALKywwr982pQaROL6rBwAK35wCAv6ZsdIBexLzyDVKlm9",
	})

	jsonFile, _ := os.Open(config.NotesJsonFilePath)
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var notes []map[string]interface{}

	err := json.Unmarshal(byteValue, &notes)
	if err != nil {
		return
	}

	_, err = client.Index("notes").AddDocuments(notes)
	if err != nil {
		panic(err)
	}
}
