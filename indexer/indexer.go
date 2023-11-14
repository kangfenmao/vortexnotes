package indexer

import (
	"fmt"
	"vortex-notes/indexer/drivers"
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
	err := sqlite.InitializeDatabase()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Database initialized successfully!")

	localIndexer := drivers.LocalIndexer{}
	StartIndexer(localIndexer)
}
