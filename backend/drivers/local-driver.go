package drivers

import (
	"encoding/json"
	"errors"
	stripmd "github.com/writeas/go-strip-markdown"
	"io"
	"os"
	"vortexnotes/backend/config"
	"vortexnotes/backend/database"
	"vortexnotes/backend/logger"
	"vortexnotes/backend/types"
)

type LocalDriver struct {
}

func (local LocalDriver) ListNotes() []string {
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

func (local LocalDriver) ParseNote(content string) string {
	return stripmd.Strip(content)
}

func (local LocalDriver) NoteExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo != nil
}

func (local LocalDriver) CreateNote(title string, content string) (err error, note types.NoteDocument) {
	filePath := config.LocalNotePath + title + ".md"
	note = types.NoteDocument{
		ID:      "",
		Name:    title,
		Content: content,
	}

	if local.NoteExist(filePath) {
		return errors.New("note file already exists"), note
	}

	file, err := os.Create(filePath)
	if err != nil {
		logger.Logger.Println("Create file error: ", err)
		return err, note
	}

	_, err = file.WriteString(content)
	if err != nil {
		logger.Logger.Println("Write file error: ", err)
		return err, note
	}
	defer file.Close()

	dbErr, noteModal := local.AddNoteToDatabase(filePath)
	if dbErr != nil {
		logger.Logger.Println("Add note to database: ", err)
		return dbErr, note
	}

	note.ID = noteModal.ID

	err = local.AddNoteToMeiliSearch(note)
	if err != nil {
		logger.Logger.Println("Add note to MeiliSearch error: ", err)
		return err, note
	}

	return nil, note
}

func (local LocalDriver) DeleteNote(id string) error {
	var note database.Note
	var result = database.DB.First(&note, "id = ?", id)

	if result.RowsAffected == 0 {
		return nil
	}

	// Delete file
	filePath := config.LocalNotePath + note.Name
	err := os.Remove(filePath)
	if err != nil {
		logger.Logger.Println("Error deleting file:", err)
		return err
	}

	// Delete Index
	index := config.MeiliSearchClient.Index("notes")
	_, err = index.DeleteDocument(note.ID)
	if err != nil {
		logger.Logger.Println("Error deleting index:", err)
		return err
	}

	// Delete db record
	database.DB.Delete(&note)

	return nil
}

func (local LocalDriver) AddNoteToDatabase(path string) (err error, note database.Note) {
	id, _ := CalculateFileHash(path)
	note = database.Note{ID: "", Name: "", Content: ""}

	fileInfo, err := os.Stat(path)
	if err != nil {
		logger.Logger.Println("Stat File Error:", err)
		return err, note
	}

	content, err := os.ReadFile(path)
	if err != nil {
		logger.Logger.Println("ReadFile Error:", err)
		return err, note
	}

	note = database.Note{ID: id, Name: fileInfo.Name(), Content: local.ParseNote(string(content))}
	result := database.DB.FirstOrCreate(&note)
	if !errors.Is(err, result.Error) {
		logger.Logger.Println("CreateNote Error:", err)
		return err, note
	}

	return nil, note
}

func (local LocalDriver) GenerateNotesJsonFile() error {
	var notes []database.Note
	result := database.DB.Find(&notes)

	if result.Error != nil {
		logger.Logger.Println("List all notes error", result.Error)
		return result.Error
	}

	jsonData, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		logger.Logger.Println("Error encoding json:", err)
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

func (local LocalDriver) AddNoteToMeiliSearch(note types.NoteDocument) error {
	documents := []map[string]interface{}{
		{
			"id":      note.ID,
			"name":    note.Name,
			"content": note.Content,
		},
	}

	_, err := config.MeiliSearchClient.Index("notes").UpdateDocuments(documents)
	if err != nil {
		return err
	}

	return nil
}

func (local LocalDriver) AddNotesToMeiliSearch() error {
	jsonFile, _ := os.Open(config.NotesJsonFilePath)
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var notes []map[string]interface{}

	err := json.Unmarshal(byteValue, &notes)
	if err != nil {
		return err
	}

	notesIndex := config.MeiliSearchClient.Index("notes")

	_, err = notesIndex.AddDocuments(notes)
	if err != nil {
		return err
	}

	_, err = notesIndex.UpdateDistinctAttribute("ID")
	if err != nil {
		return err
	}

	return nil
}
