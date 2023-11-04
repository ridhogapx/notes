package model

import (
	"time"

	"gorm.io/gorm"
)

type Notes struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Body      string
	IsDone    bool
	CreatedAt time.Time
}

type notesModelImpl struct {
  DB *gorm.DB 
}

func (db *notesModelImpl) CreateNote() {
  
}
