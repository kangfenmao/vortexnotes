package drivers

import (
	"vortex-notes/indexer/logger"
	"vortex-notes/indexer/sqlite"
	"vortex-notes/indexer/utils/file"
)

type LocalIndexer struct {
}

func (local LocalIndexer) ListAllNotes() []string {
	const path = "./data/vortexnotes"

	err := file.CreateDirectoryIfNotExists(path)
	if err != nil {
		var list []string
		logger.Logger.Fatal("Error:", err)
		return list
	}

	var notes, _ = file.ListTextFiles(path)
	return notes
}

func (local LocalIndexer) IndexExist(path string) bool {
	return false
}

func (local LocalIndexer) ParseNote(path string) string {
	return ""
}

func (local LocalIndexer) AddNoteToIndex(path string, note string) bool {
	err := sqlite.InsertFile(path)
	if err != nil {
		return false
	}

	return true
}
