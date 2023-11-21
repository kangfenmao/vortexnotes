package types

import "vortexnotes/app/database"

type Driver interface {
	ListNotes() []string
	ParseNote(content string) string
	NoteExist(path string) bool
	CreateNote(title string, content string) (err error, note NoteDocument)
	AddNoteToDatabase(path string) (err error, note database.Note)
	GenerateNotesJsonFile() error
	AddNotesToMeiliSearch() error
}
