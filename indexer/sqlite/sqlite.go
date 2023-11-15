package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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
		CREATE TABLE IF NOT EXISTS notes (
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

func InsertNote(id string, name string, content []byte) error {
	logger.Logger.Println("Insert note to db:", id, name)

	err := InsertOrUpdateNote(id, name, content)
	if err != nil {
		return err
	}

	return err
}

func InsertOrUpdateNote(id string, name string, content []byte) error {
	stmt, err := db.Prepare("INSERT OR REPLACE INTO notes(id, name, content) VALUES(?, ?, ?)")
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
