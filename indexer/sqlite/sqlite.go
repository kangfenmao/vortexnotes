package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"vortex-notes/indexer/logger"
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
			id TEXT PRIMARY KEY unique,
			name TEXT,
			content TEXT
		);
	`)

	if err != nil {
		return err
	}

	logger.Logger.Println("Database initialized successfully!")

	return nil
}

func InsertFile(filePath string, id string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.Logger.Fatal("ReadFile Error:", err)
		return err
	}

	logger.Logger.Println("Add file to db:", fileInfo.Name(), id)

	err = InsertOrUpdateFile(id, fileInfo.Name(), content)
	if err != nil {
		return err
	}

	return err
}

func InsertOrUpdateFile(id string, name string, content []byte) error {
	stmt, err := db.Prepare("INSERT OR REPLACE INTO files(id, name, content) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logger.Logger.Println("Close connection error:", err)
		}
	}(stmt)

	_, err = stmt.Exec(id, name, content)
	if err != nil {
		return err
	}

	return nil
}
