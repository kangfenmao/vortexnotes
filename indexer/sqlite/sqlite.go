package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func InitializeDatabase() error {
	dbFilePath := "./data/vortex.db"

	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		db, err := sql.Open("sqlite3", dbFilePath)
		if err != nil {
			return err
		}
		defer db.Close()

		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS files (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				filename TEXT,
				filepath TEXT,
				filehash TEXT
			)
		`)
		if err != nil {
			return err
		}

		fmt.Println("Database file created successfully!")
	} else if err != nil {
		return err
	}

	return nil
}
