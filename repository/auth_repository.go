package repository

import "notes/model"

type AuthRepository interface {
	CreateUser(*model.User) error 
  FindUser(email string) (*model.User, error)
}
