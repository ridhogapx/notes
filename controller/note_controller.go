package controller

import (
	"fmt"
	"net/http"
	"notes-v1/model"
	"os/exec"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type NoteController struct {
  Model model.NoteModel
}

type WebModel struct {
  Title string `json:"title"`
  Body string `json:"body"`
  IsDone bool `json:"is_done"`
}

func NewNoteController(model model.NoteModel) *NoteController {
  return &NoteController{
    Model: model,
  }
}

func (controller *NoteController) Route(app *gin.Engine) {
  app.Use(static.Serve("/", static.LocalFile("view/", true)))
  app.GET("/api/foo", controller.Ping)
  app.POST("/api/v1/note", controller.AddNote)
}

func (controller *NoteController) Ping(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{
    "message": "Bar",
  })
}

func (controller *NoteController) AddNote(ctx *gin.Context) {
   var data WebModel 

   ctx.ShouldBind(&data)
  
  // Generate Unique ID 
  id, _ := exec.Command("uuidgen").Output()
  
   err := controller.Model.CreateNote(&model.Note{
    ID: string(id),
    Title: data.Title,
    Body: data.Body,
    IsDone: false,
    CreatedAt: time.Now(),
   })

   if err != nil {
      fmt.Println("Failed to create note:", err)
      ctx.JSON(http.StatusInternalServerError, gin.H{
        "status" : "failure",
        "message" : "Failed to create note because internal server error",  
      })
      return
   }

   ctx.JSON(http.StatusCreated, gin.H{
     "status" : "success",
     "message" : "Successfully adding note",
   })
}
