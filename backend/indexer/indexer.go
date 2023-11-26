package indexer

import (
	"vortexnotes/backend/indexer/local"
	"vortexnotes/backend/logger"
)

func Start() {
	StartIndex(local.Driver{})
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

	logger.Logger.Println("Indexer AddNotesToDatabase")
	for _, note := range notes {
		_, err := driver.AddNoteToDatabase(note)
		if err != nil {
			return
		}
	}

	logger.Logger.Println("Indexer AddNotesToIndex")
	err = driver.AddNotesToIndex()
	if err != nil {
		logger.Logger.Fatal("Add notes to index error ", err)
		return
	}

	logger.Logger.Println("Indexer done.")
}
