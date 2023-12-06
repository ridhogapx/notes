package repository

import "gorm.io/gorm"

func NewRepository(db *gorm.DB) (AuthRepository, NoteRepository) {
	repos := &repositoryImpl{
		q: db,
	}

	return repos, repos
}
