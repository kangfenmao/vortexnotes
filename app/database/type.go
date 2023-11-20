package database

type Note struct {
	ID      string `gorm:"primaryKey" json:"hash"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
