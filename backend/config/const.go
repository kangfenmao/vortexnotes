package config

import (
	"os"
	"path/filepath"
)

var exePath, _ = os.Executable()
var basePath = filepath.Dir(exePath)
var LocalNotePath = basePath + "/data/notes/"
var AppDataPath = basePath + "/data/vortexnotes/"
var AppDbPath = AppDataPath + "app.db"
var IndexPath = AppDataPath + "notes.bleve"

const ApiHost = "0.0.0.0"
const ApiPort = "10060"
