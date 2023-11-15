package indexer

import (
	"vortex-notes/indexer/drivers/local"
	"vortex-notes/indexer/interfaces"
	"vortex-notes/indexer/logger"
	"vortex-notes/indexer/sqlite"
)

func Start() {
	err := sqlite.InitializeDatabase()
	if err != nil {
		logger.Logger.Fatal("InitializeDatabase Error:", err)
		return
	}

	StartIndex(local.Driver{})
}

func StartIndex(driver interfaces.StorageDriver) {
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
