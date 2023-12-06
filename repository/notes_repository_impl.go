package repository

import (
	"notes/model"
)

func (repos *repositoryImpl) CreateNote(note *model.Note) error {
	exec := repos.q.Create(note)

	if exec.Error != nil {
		return exec.Error
	}

	return nil
}

func (repos *repositoryImpl) FindNote(id string) *model.Note {
	var data model.Note

	exec := repos.q.Where("id = ?", id).First(&data)

	if exec.Error != nil {
		return nil
	}

	return &data

}
