package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"vortexnotes/indexer/config"
	"vortexnotes/indexer/logger"
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

	return nil
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

func ListAllNotes() (err error, allNotes []Note) {
	var notes []Note

	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		logger.Logger.Println("Error querying database:", err)
		return err, notes
	}
	defer rows.Close()

	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Name, &note.Content)
		if err != nil {
			logger.Logger.Println("Error scanning row:", err)
			return err, notes
		}
		notes = append(notes, note)
	}

	err = rows.Err()
	if err != nil {
		logger.Logger.Println("Error iterating over rows:", err)
		return err, notes
	}

	return err, notes
}
