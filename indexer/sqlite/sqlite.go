package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"vortex-notes/indexer/utils"
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
			hash TEXT
		);
	`)

	if err != nil {
		return err
	}

	return nil
}

func InsertFile(filePath string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	hash, _ := utils.CalculateFileHash(filePath)

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Println(fileInfo.Name(), hash)

	stmt, err := db.Prepare("INSERT INTO files(name, content, hash) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Close connection error:", err)
		}
	}(stmt)

	_, err = stmt.Exec(fileInfo.Name(), content, hash)
	if err != nil {
		return err
	}

	return nil
}
