package local

import (
	"os"
	"vortex-notes/indexer/config"
	"vortex-notes/indexer/logger"
	"vortex-notes/indexer/sqlite"
)

type Driver struct {
}

func (local Driver) ListNotes() []string {
	err := CreateDirectoryIfNotExists(config.LocalNotePath)
	if err != nil {
		var list []string
		logger.Logger.Fatal("Error:", err)
		return list
	}

	var notes, _ = ListTextFiles(config.LocalNotePath)
	return notes
}

func (local Driver) AddNote(path string) bool {
	id, _ := CalculateFileHash(path)

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	content, err := os.ReadFile(path)
	if err != nil {
		logger.Logger.Fatal("ReadFile Error:", err)
		return false
	}

	err = sqlite.InsertNote(id, fileInfo.Name(), content)
	if err != nil {
		return false
	}

	return true
}
