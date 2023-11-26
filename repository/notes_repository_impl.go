package repository

import "gorm.io/gorm"

type notesRepositoryImpl struct {
	q *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &notesRepositoryImpl{
		q: db,
	}
}
