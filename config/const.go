package config

import "github.com/meilisearch/meilisearch-go"

const LocalNotePath = "data/notes/"
const AppDataPath = "data/vortexnotes/"
const AppDbPath = AppDataPath + "app.db"
const NotesJsonFilePath = AppDataPath + "notes.json"
const MeiliSearchHost = "http://localhost:7700"

var MeiliSearchClient = meilisearch.NewClient(meilisearch.ClientConfig{
	Host: MeiliSearchHost,
})
