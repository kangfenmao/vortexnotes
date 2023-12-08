package main

import (
	"vortexnotes/backend/api"
	"vortexnotes/backend/database"
	"vortexnotes/backend/indexer"
	"vortexnotes/backend/storages"
)

func main() {
	storages.Init()
	database.Init()
	indexer.Start()
	api.Start()
}
