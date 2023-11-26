package repository

import (
	"notes/model"

	"gorm.io/gorm"
)

type notesRepositoryImpl struct {
	q *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &notesRepositoryImpl{
		q: db,
	}
}

func (repos *notesRepositoryImpl) CreateNote(note *model.Notes) error {
	exec := repos.q.Create(note)

	if exec.Error != nil {
		return exec.Error
	}

	return nil
}
