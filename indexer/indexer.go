package indexer

import (
	"vortex-notes/indexer/drivers"
	"vortex-notes/indexer/drivers/local"
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

func StartIndex(driver drivers.StorageDriver) {
	notes := driver.ListNotes()
	for _, note := range notes {
		driver.AddIndex(driver.AddNote(note))
	}
}
