package local

import (
	"errors"
	stripmd "github.com/writeas/go-strip-markdown"
	"os"
	"vortexnotes/backend/blevesearch"
	"vortexnotes/backend/config"
	"vortexnotes/backend/database"
	"vortexnotes/backend/logger"
	"vortexnotes/backend/types"
)

type Driver struct{}

func (localDriver Driver) BeforeStart() error {
	database.DB.Where("1 = 1").Delete(&database.Note{})
	blevesearch.ResetIndex()
	return nil
}

func (localDriver Driver) ListNotes() []string {
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

func (localDriver Driver) ParseNote(content string) string {
	return stripmd.Strip(content)
}

func (localDriver Driver) NoteExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo != nil
}

func (localDriver Driver) CreateNote(title string, content string) (note types.NoteDocument, err error) {
	filePath := config.LocalNotePath + title + ".md"
	note = types.NoteDocument{
		ID:      "",
		Name:    title,
		Content: content,
	}

	if localDriver.NoteExist(filePath) {
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

	noteModal, dbErr := localDriver.AddNoteToDatabase(filePath)
	if dbErr != nil {
		logger.Logger.Println("Add note to database: ", err)
		return note, dbErr
	}

	note.ID = noteModal.ID

	err = localDriver.AddNoteToIndex(note)
	if err != nil {
		logger.Logger.Println("Add note to index error: ", err)
		return note, err
	}

	return note, nil
}

func (localDriver Driver) DeleteNote(id string) error {
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

func (localDriver Driver) AddNoteToDatabase(path string) (note database.Note, err error) {
	id, _ := CalculateFileHash(path)
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

	note = database.Note{ID: id, Name: fileInfo.Name(), Content: localDriver.ParseNote(string(content))}
	result := database.DB.CreateInBatches(&note, 100)
	if !errors.Is(err, result.Error) {
		logger.Logger.Println("CreateNote Error:", err)
		return note, err
	}

	return note, nil
}

func (localDriver Driver) AddNoteToIndex(note types.NoteDocument) error {
	err := blevesearch.NotesIndex.Index(note.ID, note)
	if err != nil {
		return err
	}

	return nil
}

func (localDriver Driver) AddNotesToIndex() error {
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
