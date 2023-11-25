package indexer

import (
	"vortexnotes/backend/logger"
)

func Start() {
	StartIndex(LocalIndexer{})
}

func StartIndex(driver IndexerDriver) {
	logger.Logger.Println("Indexer start")

	notes := driver.ListNotes()

	for _, note := range notes {
		err, _ := driver.AddNoteToDatabase(note)
		if err != nil {
			return
		}
	}

	err := driver.AddNotesToMeiliSearch()
	if err != nil {
		logger.Logger.Println("Add notes to meilisearch error", err)
		return
	}

	logger.Logger.Println("Indexer done.")
}
