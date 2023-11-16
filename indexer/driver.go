package indexer

type Driver interface {
	ListNotes() []string
	ParseNote(content string) string
	AddNoteToDatabase(path string)
	GenerateNotesJsonFile() error
	AddNotesToMeiliSearch() error
}
