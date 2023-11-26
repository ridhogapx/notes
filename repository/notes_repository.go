package repository

import "notes/model"

type NoteRepository interface {
	CreateNote(*model.Notes) error
	FindNote(id string) *model.Notes
}
