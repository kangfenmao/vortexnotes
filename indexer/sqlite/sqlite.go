package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/goccy/go-json"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"vortex-notes/indexer/config"
	"vortex-notes/indexer/logger"
)

var db *sql.DB

type Note struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func InitializeDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", config.AppDbPath)
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

func GenerateNotesJsonFile() {
	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		logger.Logger.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	var notes []Note

	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Name, &note.Content)
		if err != nil {
			logger.Logger.Println("Error scanning row:", err)
			return
		}
		notes = append(notes, note)
	}

	err = rows.Err()
	if err != nil {
		logger.Logger.Println("Error iterating over rows:", err)
		return
	}

	jsonData, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		logger.Logger.Println("Error encoding JSON:", err)
		return
	}

	file, err := os.Create(config.NotesJsonFilePath)
	if err != nil {
		logger.Logger.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		logger.Logger.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON file created successfully.")
}
