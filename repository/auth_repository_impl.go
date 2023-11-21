package repository

import (
	"notes/model"

	"gorm.io/gorm"
)

type authRepositoryImpl struct {
  q *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
  return &authRepositoryImpl{
    q: db,
  }
}

func (repo *authRepositoryImpl) CreateUser(user *model.User) error {
  exec := repo.q.Create(user)
  
  if exec.Error != nil {
    return exec.Error
  }

  return nil
}

func (repo *authRepositoryImpl) FindUser(email string) (*model.User, error) {
  var user model.User

  exec := repo.q.Where("email = ?", email).First(&user)
  
  if exec.Error != nil {
    return nil, exec.Error
  }

  return &user, nil
}
