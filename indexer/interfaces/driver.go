package interfaces

type StorageDriver interface {
	ListNotes() []string
	ParseNote(content []byte) []byte
	AddNoteToDatabase(path string)
	GenerateNotesJsonFile() error
	AddNotesToMeiliSearch() error
}
