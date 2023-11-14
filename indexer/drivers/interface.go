package drivers

type StorageDriver interface {
	ListNotes() []string
	AddNote(path string) bool
}
