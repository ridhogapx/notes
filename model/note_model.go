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

type NoteModel interface {
  CreateNote(*Notes) error
  FindNote(id string) (*Notes, error)
}

type notesModelImpl struct {
  DB *gorm.DB 
}

