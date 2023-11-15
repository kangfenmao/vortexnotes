package local

import (
	"github.com/goccy/go-json"
	stripmd "github.com/writeas/go-strip-markdown"
	"io"
	"os"
	"vortexnotes/config"
	"vortexnotes/indexer/logger"
	"vortexnotes/indexer/sqlite"
)

type Driver struct {
}

func (local Driver) ListNotes() []string {
	var notes []string

	err := CreateDirectoryIfNotExists(config.LocalNotePath)
	if err != nil {
		logger.Logger.Fatal("Error:", err)
		return notes
	}

	err, notes = ListTextFiles(config.LocalNotePath)
	if err != nil {
		logger.Logger.Println("List text files error", err)
		return notes
	}

	return notes
}

func (local Driver) ParseNote(content string) string {
	return stripmd.Strip(content)
}

func (local Driver) AddNoteToDatabase(path string) {
	id, _ := CalculateFileHash(path)

	fileInfo, err := os.Stat(path)
	if err != nil {
		logger.Logger.Println("Stat File Error:", err)
		return
	}

	content, err := os.ReadFile(path)
	if err != nil {
		logger.Logger.Println("ReadFile Error:", err)
		return
	}

	err = sqlite.InsertNote(id, fileInfo.Name(), local.ParseNote(string(content)))
	if err != nil {
		logger.Logger.Println("InsertNote Error:", err)
		return
	}

	return
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

	return nil
}

func (local Driver) AddNotesToMeiliSearch() error {
	jsonFile, _ := os.Open(config.NotesJsonFilePath)
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var notes []map[string]interface{}

	err := json.Unmarshal(byteValue, &notes)
	if err != nil {
		return err
	}

	_, err = config.MeiliSearchClient.Index("notes").AddDocuments(notes)
	if err != nil {
		return err
	}

	return nil
}
