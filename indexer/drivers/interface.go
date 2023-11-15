package drivers

type StorageDriver interface {
	ListNotes() []string
	AddNote(path string) (string, string, string)
	AddIndex(id string, name string, content string)
}
