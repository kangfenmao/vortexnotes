package main

import (
	"vortexnotes/app/api"
	"vortexnotes/app/indexer"
)

func main() {
	indexer.Start()
	api.Start()
}
