package drivers

import (
	"fmt"
	"vortex-notes/indexer/sqlite"
	"vortex-notes/indexer/utils"
)

type LocalIndexer struct {
}

func (local LocalIndexer) ListAllNotes() []string {
	const path = "./data/vortexnotes"

	err := utils.CreateDirectoryIfNotExists(path)
	if err != nil {
		var list []string
		fmt.Println("Error:", err)
		return list
	}

	var notes, _ = utils.ListTextFiles(path)
	return notes
}

func (local LocalIndexer) IndexExist(path string) bool {
	return false
}

func (local LocalIndexer) ParseNote(path string) string {
	return ""
}

func (local LocalIndexer) AddNoteToIndex(path string, note string) bool {
	hash, err := utils.CalculateFileHash(path)

	if err != nil {
		return true
	}

	err = sqlite.InsertFile(path)
	if err != nil {
		return false
	}

	fmt.Println("AddNoteToIndex: ", path, hash)

	return true
}
