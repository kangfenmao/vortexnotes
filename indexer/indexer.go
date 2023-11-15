package indexer

import (
	"vortex-notes/indexer/drivers/local"
	"vortex-notes/indexer/interfaces"
	"vortex-notes/indexer/logger"
	"vortex-notes/indexer/sqlite"
)

func Start() {
	logger.Logger.Println("Indexer start")

	err := sqlite.InitializeDatabase()
	if err != nil {
		logger.Logger.Fatal("InitializeDatabase Error:", err)
		return
	}

	StartIndex(local.Driver{})
}

func StartIndex(driver interfaces.StorageDriver) {
	notes := driver.ListNotes()

	for _, note := range notes {
		err := driver.AddNoteToDatabase(note)
		if err != nil {
			logger.Logger.Println("Add notes to database error", err)
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
