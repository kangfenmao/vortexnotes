package indexer

import (
	"errors"
	stripmd "github.com/writeas/go-strip-markdown"
	"os"
	"vortexnotes/backend/config"
	"vortexnotes/backend/database"
	"vortexnotes/backend/logger"
	"vortexnotes/backend/storage"
	"vortexnotes/backend/types"
)

type LocalIndexer struct {
}

func (localIndexer LocalIndexer) BeforeStart() error {
	database.DB.Where("1 = 1").Delete(&database.Note{})

	_, err := config.MeiliSearchClient.DeleteIndex("notes")
	if err != nil {
		return err
	}

	return nil
}

func (localIndexer LocalIndexer) ListNotes() []string {
	var notes []string

	err := storage.CreateDirectoryIfNotExists(config.LocalNotePath)
	if err != nil {
		logger.Logger.Fatal("Error:", err)
		return notes
	}

	err, notes = storage.ListTextFiles(config.LocalNotePath)
	if err != nil {
		logger.Logger.Println("List text files error", err)
		return notes
	}

	return notes
}

func (localIndexer LocalIndexer) ParseNote(content string) string {
	return stripmd.Strip(content)
}

func (localIndexer LocalIndexer) NoteExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo != nil
}

func (localIndexer LocalIndexer) CreateNote(title string, content string) (err error, note types.NoteDocument) {
	filePath := config.LocalNotePath + title + ".md"
	note = types.NoteDocument{
		ID:      "",
		Name:    title,
		Content: content,
	}

	if localIndexer.NoteExist(filePath) {
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

	dbErr, noteModal := localIndexer.AddNoteToDatabase(filePath)
	if dbErr != nil {
		logger.Logger.Println("Add note to database: ", err)
		return dbErr, note
	}

	note.ID = noteModal.ID

	err = localIndexer.AddNoteToMeiliSearch(note)
	if err != nil {
		logger.Logger.Println("Add note to MeiliSearch error: ", err)
		return err, note
	}

	return nil, note
}

func (localIndexer LocalIndexer) DeleteNote(id string) error {
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

func (localIndexer LocalIndexer) AddNoteToDatabase(path string) (err error, note database.Note) {
	id, _ := storage.CalculateFileHash(path)
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

	note = database.Note{ID: id, Name: fileInfo.Name(), Content: localIndexer.ParseNote(string(content))}
	result := database.DB.CreateInBatches(&note, 100)
	if !errors.Is(err, result.Error) {
		logger.Logger.Println("CreateNote Error:", err)
		return err, note
	}

	return nil, note
}

func (localIndexer LocalIndexer) AddNoteToMeiliSearch(note types.NoteDocument) error {
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

func (localIndexer LocalIndexer) AddNotesToMeiliSearch() error {
	var notes []database.Note
	database.DB.Find(&notes)

	_, err := config.MeiliSearchClient.DeleteIndex("notes")
	if err != nil {
		return err
	}

	notesIndex := config.MeiliSearchClient.Index("notes")

	_, err = notesIndex.AddDocuments(notes)
	if err != nil {
		return err
	}

	return nil
}
