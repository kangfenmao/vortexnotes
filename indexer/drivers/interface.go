package drivers

type StorageDriver interface {
	ListNotes() []string
	AddNoteToDatabase(path string) error
	GenerateNotesJsonFile() error
	AddNotesToMeiliSearch() error
}
