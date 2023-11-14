package indexer

import (
	"vortex-notes/indexer/drivers"
)

func Start() {
	localIndexer := drivers.LocalIndexer{}
	StartIndexer(localIndexer)
}
