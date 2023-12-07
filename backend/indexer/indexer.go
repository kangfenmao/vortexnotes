package indexer

import (
	"errors"
	"os"
	"vortexnotes/backend/blevesearch"
	"vortexnotes/backend/config"
	"vortexnotes/backend/database"
	"vortexnotes/backend/logger"
	"vortexnotes/backend/storages"
	"vortexnotes/backend/types"
)

func Start() {
	StartIndex()
}

func StartIndex() {
	logger.Logger.Println("Indexer start")

	ResetIndex()

	logger.Logger.Println("Indexer ListNotes")
	notes := ListNotes()

	logger.Logger.Println("Indexer AddNotesToDatabase")
	for _, note := range notes {
		_, err := AddNoteToDatabase(note)
		if err != nil {
			return
		}
	}

	logger.Logger.Println("Indexer AddNotesToIndex")
	err := AddNotesToIndex()
	if err != nil {
		logger.Logger.Fatal("Add notes to index error ", err)
		return
	}

	logger.Logger.Println("Indexer done.")
}

func ResetIndex() {
	database.DB.Where("1 = 1").Delete(&database.Note{})
	blevesearch.ResetIndex()
}

func ListNotes() []string {
	var notes []string

	err := storages.CreateDirectoryIfNotExists(config.LocalNotePath)
	if err != nil {
		logger.Logger.Fatal("Error:", err)
		return notes
	}

	err, notes = storages.ListTextFiles(config.LocalNotePath)
	if err != nil {
		logger.Logger.Println("List text files error", err)
		return notes
	}

	return notes
}

func NoteExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo != nil
}

func CreateNote(title string, content string) (note types.NoteDocument, err error) {
	filePath := config.LocalNotePath + title + ".md"
	note = types.NoteDocument{
		ID:      "",
		Name:    title,
		Content: content,
	}

	if NoteExist(filePath) {
		return note, errors.New("note file already exists")
	}

	file, err := os.Create(filePath)
	if err != nil {
		logger.Logger.Println("Create file error: ", err)
		return note, err
	}

	_, err = file.WriteString(content)
	if err != nil {
		logger.Logger.Println("Write file error: ", err)
		return note, err
	}
	defer file.Close()

	noteModal, dbErr := AddNoteToDatabase(filePath)
	if dbErr != nil {
		logger.Logger.Println("Add note to database: ", err)
		return note, dbErr
	}

	note.ID = noteModal.ID

	err = AddNoteToIndex(note)
	if err != nil {
		logger.Logger.Println("Add note to index error: ", err)
		return note, err
	}

	return note, nil
}

func DeleteNote(id string) error {
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
	err = blevesearch.NotesIndex.Delete(note.ID)
	if err != nil {
		logger.Logger.Println("Error deleting index:", err)
		return err
	}

	// Delete db record
	database.DB.Delete(&note)

	return nil
}

func AddNoteToDatabase(path string) (note database.Note, err error) {
	id, _ := storages.CalculateFileHash(path)
	note = database.Note{ID: "", Name: "", Content: ""}

	fileInfo, err := os.Stat(path)
	if err != nil {
		logger.Logger.Println("Stat File Error:", err)
		return note, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		logger.Logger.Println("ReadFile Error:", err)
		return note, err
	}

	note = database.Note{
		ID:        id,
		Name:      fileInfo.Name(),
		Content:   string(content),
		CreatedAt: fileInfo.ModTime(),
		UpdatedAt: fileInfo.ModTime(),
	}

	result := database.DB.CreateInBatches(&note, 100)
	if !errors.Is(err, result.Error) {
		logger.Logger.Println("CreateNote Error:", err)
		return note, err
	}

	return note, nil
}

func AddNoteToIndex(note types.NoteDocument) error {
	err := blevesearch.NotesIndex.Index(note.ID, note)
	if err != nil {
		return err
	}

	return nil
}

func AddNotesToIndex() error {
	var notes []database.Note
	database.DB.Find(&notes)

	batch := blevesearch.NotesIndex.NewBatch()
	for _, note := range notes {
		batch.Index(note.ID, note)
	}

	err := blevesearch.NotesIndex.Batch(batch)
	if err != nil {
		logger.Logger.Println("Failed to perform batch index operation:", err)
		return err
	}

	return nil
}
