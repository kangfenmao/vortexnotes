package drivers

import "fmt"

type LocalIndexer struct {
}

func (local LocalIndexer) ListAllNotes() []string {
	var notes []string
	return notes
}

func (local LocalIndexer) IndexExist(path string) bool {
	fmt.Println("IndexExist: ", path)
	return false
}

func (local LocalIndexer) ParseNote(path string) string {
	fmt.Println("ParseNote: ", path)
	return ""
}

func (local LocalIndexer) AddNoteToIndex(path string, note string) bool {
	fmt.Println("AddNoteToIndex: ", path)
	return true
}
