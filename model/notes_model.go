package model

import "time"

type Notes struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Body      string
	IsDone    bool
	CreatedAt time.Time
}
