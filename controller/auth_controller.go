package controller

import (
	"fmt"
	"net/http"
	"notes/model"
	"notes/repository"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Repos repository.AuthRepository
}

const (
	BootstrapCSS = "https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css"
	BootstrapJS  = "https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js"
)

func NewAuthController(repos repository.AuthRepository) *AuthController {
	return &AuthController{
		Repos: repos,
	}
}

func (controller *AuthController) Routes(r *gin.Engine) {
	r.GET("/signup", controller.SignUpView)
}

func (controller *AuthController) SignUpView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{
		"title":  "Sign Up - Notes",
		"style":  BootstrapCSS,
		"script": BootstrapJS,
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

}
