package drivers

type StorageDriver interface {
	ListNotes() []string
	AddNoteToDatabase(path string)
	SyncNoteToMeiliSearch()
}
