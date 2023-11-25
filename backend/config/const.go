package config

import (
	"github.com/meilisearch/meilisearch-go"
	"vortexnotes/backend/utils"
)

const LocalNotePath = "data/notes/"
const AppDataPath = "data/vortexnotes/"
const AppDbPath = AppDataPath + "app.db"
const ApiHost = "0.0.0.0"
const ApiPort = "7701"

var MeiliSearchClient = meilisearch.NewClient(meilisearch.ClientConfig{
	Host: utils.MeiliSearchHost(),
})
