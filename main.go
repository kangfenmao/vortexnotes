package main

import (
	"vortexnotes/app/api"
	"vortexnotes/app/database"
	"vortexnotes/app/indexer"
)

func main() {
	database.Init()
	go indexer.Start()
	api.Start()
}
