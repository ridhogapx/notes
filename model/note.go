package model

type Note struct {
	ID     string `gorm:"primaryKey"`
	Title  string
	Body   string
	Author string
}
