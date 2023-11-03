package controller

import "github.com/gin-gonic/gin"

type notesController struct {
	R *gin.Engine
}

func NewNotesController(r *gin.Engine) *notesController {
  return &notesController{
    R: r,
  }
}

func (controller *notesController) Home() {
  controller.R.Static("/", "./view")
}
