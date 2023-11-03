package main

import (
	"notes-v1/controller"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  
  // Inject controller dependency
  notesController := controller.NewNotesController(r)
  notesController.Init()

  r.Run(":8080")
}
