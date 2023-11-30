package main

import (
	"fmt"
	"notes/config"
	"notes/controller"
	"notes/repository"
	"os"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env")

  if err != nil {
    fmt.Println("Failed to load .env")
    return
  }

  r := gin.Default()

  DB_SOURCE := os.Getenv("DB_SOURCE")
  SECRET := os.Getenv("SECRET")

  store := cookie.NewStore([]byte(SECRET))

  dbConn := config.NewDBConnection(DB_SOURCE)
  authRepos := repository.NewAuthRepository(dbConn)
  authController := controller.NewAuthController(authRepos)

  r.LoadHTMLGlob("views/*")

  authController.SetupSession(store, r) 
  authController.Routes(r)


  r.Run(":3000")
}
