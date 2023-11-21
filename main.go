package main

import (
	"vortexnotes/app/api"
	"vortexnotes/app/database"
	"vortexnotes/app/indexer"
)

func main() {
	database.InitializeDatabase()
	indexer.Start()
	api.Start()
}
