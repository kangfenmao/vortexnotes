package indexer

import (
	"vortexnotes/app/database"
	"vortexnotes/indexer/drivers/local"
	"vortexnotes/indexer/logger"
)

func Start() {
	database.InitializeDatabase()
	StartIndex(local.Driver{})
}

func StartIndex(driver Driver) {
	logger.Logger.Println("Indexer start")

	notes := driver.ListNotes()

	for _, note := range notes {
		driver.AddNoteToDatabase(note)
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
