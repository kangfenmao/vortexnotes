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

func (local Driver) AddNoteToDatabase(path string) error {
	id, _ := CalculateFileHash(path)

	fileInfo, err := os.Stat(path)
	if err != nil {
		logger.Logger.Println("Stat File Error:", err)
		return err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		logger.Logger.Println("ReadFile Error:", err)
		return err
	}

	err = sqlite.InsertNote(id, fileInfo.Name(), content)
	if err != nil {
		logger.Logger.Println("InsertNote Error:", err)
		return err
	}

	return err
}

func (local Driver) GenerateNotesJsonFile() error {
	err, notes := sqlite.ListAllNotes()
	if err != nil {
		logger.Logger.Println("List all notes error", err)
		return err
	}

	jsonData, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		logger.Logger.Println("Error encoding JSON:", err)
		return err
	}

	file, err := os.Create(config.NotesJsonFilePath)
	if err != nil {
		logger.Logger.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		logger.Logger.Println("Error writing to file:", err)
		return err
	}

	logger.Logger.Println("JSON file created successfully.")

	return err
}

func (local Driver) AddNotesToMeiliSearch() error {
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
		return err
	}

	_, err = client.Index("notes").AddDocuments(notes)
	if err != nil {
		return err
	}

	return err
}
