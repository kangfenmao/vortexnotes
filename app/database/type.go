package database

type NoteModel struct {
	ID      string `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
