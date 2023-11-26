package main

import (
	"vortexnotes/backend/api"
	"vortexnotes/backend/database"
	"vortexnotes/backend/indexer"
)

func main() {
	database.Init()
	indexer.Start()
	api.Start()
}
