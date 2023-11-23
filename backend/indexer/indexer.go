package indexer

import (
	"vortexnotes/backend/drivers"
	"vortexnotes/backend/logger"
)

func Start() {
	StartIndex(drivers.LocalDriver{})
}

func StartIndex(driver drivers.Driver) {
	logger.Logger.Println("Indexer start")

	notes := driver.ListNotes()

	for _, note := range notes {
		err, _ := driver.AddNoteToDatabase(note)
		if err != nil {
			return
		}
	}

	err := driver.GenerateNotesJsonFile()
	if err != nil {
		logger.Logger.Println("Generate notes json file error", err)
		return
	}

	err = driver.AddNotesToMeiliSearch()
	if err != nil {
		logger.Logger.Println("Add notes to meilisearch error", err)
		return
	}

	logger.Logger.Println("Indexer done.")
}
