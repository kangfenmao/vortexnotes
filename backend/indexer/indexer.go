package indexer

import (
	"vortexnotes/backend/logger"
)

func Start() {
	StartIndex(LocalIndexer{})
}

func StartIndex(driver Driver) {
	logger.Logger.Println("Indexer start")

	err := driver.BeforeStart()
	if err != nil {
		logger.Logger.Fatal("BeforeStart index error:", err)
		return
	}

	logger.Logger.Println("Indexer ListNotes")
	notes := driver.ListNotes()

	for _, note := range notes {
		err, _ := driver.AddNoteToDatabase(note)
		if err != nil {
			return
		}
	}

	logger.Logger.Println("Indexer AddNotesToMeiliSearch")
	err = driver.AddNotesToMeiliSearch()
	if err != nil {
		logger.Logger.Fatal("Add notes to meilisearch error", err)
		return
	}

	logger.Logger.Println("Indexer done.")
}
