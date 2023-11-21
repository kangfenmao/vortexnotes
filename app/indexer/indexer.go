package indexer

import (
	"vortexnotes/app/database"
	"vortexnotes/app/drivers"
	"vortexnotes/app/logger"
	"vortexnotes/app/types"
)

func Start() {
	database.InitializeDatabase()
	StartIndex(drivers.LocalDriver{})
}

func StartIndex(driver types.Driver) {
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
