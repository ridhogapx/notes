package repository

import "gorm.io/gorm"

func NewRepository(db *gorm.DB) (AuthRepository, NoteRepository) {
	return &repositoryImpl{
		q: db,
	}
}
