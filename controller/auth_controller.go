package controller

import "notes/repository"

type AuthController struct {
  Repos repository.AuthRepository
} 

func NewAuthController(repos repository.AuthRepository) *AuthController {
  return &AuthController{
    Repos: repos,
  }
}

