package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"vortex-notes/indexer/logger"
	"vortex-notes/indexer/utils/file"
)

const DbFilePath = "./data/vortex.db"

var db *sql.DB

func InitializeDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", DbFilePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS files (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			content TEXT,
			hash TEXT unique
		);
	`)

	if err != nil {
		return err
	}

	logger.Logger.Println("Database initialized successfully!")

	return nil
}

func InsertFile(filePath string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	hash, _ := file.CalculateFileHash(filePath)

	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.Logger.Fatal("ReadFile Error:", err)
		return err
	}

	logger.Logger.Println("Add file to sqlite:", fileInfo.Name(), hash)

	err = InsertOrUpdateFile(fileInfo.Name(), content, hash)
	if err != nil {
		return err
	}

	return err
}

func InsertOrUpdateFile(name string, content []byte, hash string) error {
	stmt, err := db.Prepare("INSERT OR REPLACE INTO files(name, content, hash) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logger.Logger.Println("Close connection error:", err)
		}
	}(stmt)

	_, err = stmt.Exec(name, content, hash)
	if err != nil {
		return err
	}

	return nil
}
