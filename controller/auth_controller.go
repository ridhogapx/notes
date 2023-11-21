package controller

import (
	"net/http"
	"notes/repository"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
  Repos repository.AuthRepository
} 

func NewAuthController(repos repository.AuthRepository) *AuthController {
  return &AuthController{
    Repos: repos,
  }
}


func (controller *AuthController) Routes(r *gin.Engine) {
  r.GET("/signup", controller.SignUp) 
}

func (controller *AuthController) SignUp(ctx *gin.Context) {
  ctx.HTML(http.StatusOK, "signup.html", gin.H{
    "title" : "Sign Up - Notes",
  })
}
