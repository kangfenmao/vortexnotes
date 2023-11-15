package interfaces

type StorageDriver interface {
	ListNotes() []string
	AddNoteToDatabase(path string)
	GenerateNotesJsonFile() error
	AddNotesToMeiliSearch() error
}
