package local

import (
	"vortex-notes/indexer/logger"
	"vortex-notes/indexer/sqlite"
)

type Indexer struct {
}

func (local Indexer) ListAllNotes() []string {
	const path = "./data/vortexnotes"

	err := CreateDirectoryIfNotExists(path)
	if err != nil {
		var list []string
		logger.Logger.Fatal("Error:", err)
		return list
	}

	var notes, _ = ListTextFiles(path)
	return notes
}

func (local Indexer) IndexExist(path string) bool {
	return false
}

func (local Indexer) ParseNote(path string) string {
	return ""
}

func (local Indexer) AddNoteToIndex(path string, note string) bool {
	id, _ := CalculateFileHash(path)

	err := sqlite.InsertFile(path, id)
	if err != nil {
		return false
	}

	return true
}
