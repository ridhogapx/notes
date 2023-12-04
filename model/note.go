package model

type Note struct {
	ID     string `gorm:"primaryKey"`
	Title  string
	Body   string
	Author string
}

type NoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
