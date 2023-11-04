package controller

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type notesController struct {
	R *gin.Engine
}

func NewNotesController(r *gin.Engine) *notesController {
  return &notesController{
    R: r,
  }
}

func (controller *notesController) Init() {
  controller.R.Use(static.Serve("/", static.LocalFile("view/", true)))
  controller.R.GET("/api/foo", controller.Ping)
}

func (controller *notesController) Ping(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{
    "message": "Bar",
  })
}
