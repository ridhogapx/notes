package controller

import (
	"net/http"
	"notes-v1/model"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type NoteController struct {
  Model model.NoteModel
}

func NewNoteController(model model.NoteModel) *NoteController {
  return &NoteController{
    Model: model,
  }
}

func (controller *NoteController) Route(app *gin.Engine) {
  app.Use(static.Serve("/", static.LocalFile("view/", true)))
  app.GET("/api/foo", controller.Ping)
}

func (controller *NoteController) Ping(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{
    "message": "Bar",
  })
}
