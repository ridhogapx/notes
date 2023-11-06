package main

import (
	"log"

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

  // Inject model dependency

  // Inject controller dependency


  r.Run(":8080")
}
