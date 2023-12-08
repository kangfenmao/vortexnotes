package config

import "os"

var basePath, _ = os.Getwd()
var LocalNotePath = basePath + "/data/notes/"
var AppDataPath = basePath + "/data/vortexnotes/"
var AppDbPath = AppDataPath + "app.db"
var IndexPath = AppDataPath + "notes.bleve"

const ApiHost = "0.0.0.0"
const ApiPort = "7701"
