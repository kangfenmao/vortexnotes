package indexer

import (
	"vortex-notes/indexer/drivers"
	"vortex-notes/indexer/logger"
	"vortex-notes/indexer/sqlite"
)

type Indexer interface {
	ListAllNotes() []string
	IndexExist(path string) bool
	ParseNote(path string) string
	AddNoteToIndex(path string, note string) bool
}

func StartIndexer(indexer Indexer) {
	notes := indexer.ListAllNotes()
	for _, note := range notes {
		if !indexer.IndexExist(note) {
			parsedNote := indexer.ParseNote(note)
			indexer.AddNoteToIndex(note, parsedNote)
		}
	}
}

func Start() {
	logger.Logger.Println("Indexer start")

	err := sqlite.InitializeDatabase()
	if err != nil {
		logger.Logger.Fatal("InitializeDatabase Error:", err)
		return
	}

	localIndexer := drivers.LocalIndexer{}
	StartIndexer(localIndexer)
}
