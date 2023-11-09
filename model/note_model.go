package model

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Body      string
	IsDone    bool
	CreatedAt time.Time
}

type NoteModel interface {
  CreateNote(*Note) error
  FindNote(id string) (*[]Note, error)
}

type noteModelImpl struct {
  DB *gorm.DB 
}

func NewNoteModel(db *gorm.DB) NoteModel {
  return &noteModelImpl{
    DB: db,
  }
}

func (note *noteModelImpl) CreateNote(data *Note) error {
  exec := note.DB.Create(data)

  if exec.Error != nil {
    return exec.Error
  }

  return nil
}

func (note *noteModelImpl) FindNote(id string) (*[]Note, error) {
  var data []Note
  exec := note.DB.Where("id = ?", id).Find(&data)

  if exec.Error != nil {
    return nil, exec.Error
  }

  return &data, nil
}
