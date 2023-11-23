package drivers

import (
	"vortexnotes/backend/database"
	"vortexnotes/backend/types"
)

type Driver interface {
	ListNotes() []string
	ParseNote(content string) string
	NoteExist(path string) bool
	CreateNote(title string, content string) (err error, note types.NoteDocument)
	DeleteNote(id string) error
	AddNoteToDatabase(path string) (err error, note database.Note)
	GenerateNotesJsonFile() error
	AddNotesToMeiliSearch() error
}
