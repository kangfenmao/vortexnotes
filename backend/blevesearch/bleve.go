package blevesearch

import (
	"github.com/blevesearch/bleve/v2"
	"os"
	"vortexnotes/backend/config"
	"vortexnotes/backend/logger"
)

var NotesIndex bleve.Index

func ResetIndex() {
	deleteIndex()
	createIndex()
}

func deleteIndex() {
	_ = os.RemoveAll(config.IndexPath)
}

func createIndex() {
	var err error
	mapping := bleve.NewIndexMapping()
	NotesIndex, err = bleve.New(config.IndexPath, mapping)
	if err != nil {
		logger.Logger.Fatalln("Bleve failed to delete index:", err)
	}
}
