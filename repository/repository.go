package repository

import "gorm.io/gorm"

func NewRepository(db *gorm.DB) (AuthRepository, NoteRepository) {
	authRepos := authRepositoryImpl{
		q: db,
	}
	noteRepos := notesRepositoryImpl{
		q: db,
	}

	return &authRepos, &noteRepos
}
