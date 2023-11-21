package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"vortexnotes/app/config"
	"vortexnotes/app/logger"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.AppDbPath), &gorm.Config{})
	if err != nil {
		logger.Logger.Fatal("Initialize Database Error:", err)
		return
	}

	err = DB.AutoMigrate(&Note{})
	if err != nil {
		logger.Logger.Fatal("Migrate Database Error:", err)
		return
	}
}
