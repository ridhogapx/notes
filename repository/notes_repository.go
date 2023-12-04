package repository

import "notes/model"

type NoteRepository interface {
	CreateNote(*model.Note) error
	FindNote(id string) *model.Note
}
