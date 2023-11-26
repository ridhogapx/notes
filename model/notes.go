package model

type Notes struct {
	ID     string `gorm:"primaryKey"`
	Title  string
	Body   string
	Author string
}
