package main

import (
	"log"
	"notes-v1/config"
	"notes-v1/controller"
	"notes-v1/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
  r := gin.Default()

  // Load environment variable 
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatal("Failed to load .env")
  }

  // Load configuration database
  DB_SOURCE := os.Getenv("DB_SOURCE")
  db := config.NewDBConn(DB_SOURCE)

  // Inject model dependency
  noteModel := model.NewNoteModel(db)

  // Inject controller dependency
  noteController := controller.NewNoteController(noteModel)
  noteController.Route(r)


  r.Run(":8080")
}
