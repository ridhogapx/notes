package controller

import (
	"fmt"
	"net/http"
	"notes/helper"
	"notes/model"
	"notes/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	Repos repository.AuthRepository
}

const (
	BootstrapCSS = "https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css"
	BootstrapJS  = "https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js"
	PopperJS     = "https://cdn.jsdelivr.net/npm/@popperjs/core@2.11.8/dist/umd/popper.min.js"
)

func NewAuthController(repos repository.AuthRepository) *AuthController {
	return &AuthController{
		Repos: repos,
	}
}

func (controller *AuthController) Routes(r *gin.Engine) {
	r.GET("/signup", controller.SignUpView)
	r.GET("/", controller.NotesView)
	r.POST("/signup", controller.SignUp)
}

func (controller *AuthController) SignUpView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{
		"title":  "Sign Up - Notes",
		"style":  BootstrapCSS,
		"script": BootstrapJS,
		"popper": PopperJS,
	})
}

func (controller *AuthController) NotesView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "me.html", gin.H{
		"title": "My Notes",
	})
}

func (controller *AuthController) SignUp(ctx *gin.Context) {
	var payload model.AuthRequest

	err := ctx.ShouldBind(&payload)

	if err != nil {
		fmt.Println("Request should be in JSON format!", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Payload should be in JSON",
		})
		return
	}

	// Check if user is already registered
	user, _ := controller.Repos.FindUser(payload.Email)

	if user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "Email sudah terdaftar",
		})

		return
	}

	// If not yet, then user can register

	// We need to hash password before insert into database.
	hash := helper.HashPassword(payload.Password)

	// Generate unique ID
	id := uuid.New().String()

	err = controller.Repos.CreateUser(&model.User{
		ID:       id,
		Email:    payload.Email,
		Password: hash,
		Name:     payload.Name,
	})

	// Internal logging
	if err != nil {
		fmt.Println("Failed to create user", err)
		return
	}

	// if success, we need redirect it into Home.
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Berhasil mendaftar user baru",
	})
}
