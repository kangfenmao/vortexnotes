package indexer

import (
	"vortexnotes/backend/database"
	"vortexnotes/backend/types"
)

type Driver interface {
	BeforeStart() error
	ListNotes() []string
	ParseNote(content string) string
	NoteExist(path string) bool
	CreateNote(title string, content string) (note types.NoteDocument, err error)
	DeleteNote(id string) error
	AddNoteToDatabase(path string) (note database.Note, err error)
	AddNotesToIndex() error
}
