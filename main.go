package main

import (
	"fmt"
	"notes/config"
	"notes/controller"
	"notes/repository"
	"os"

	"github.com/gin-contrib/sessions/redis"
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

	// Redis Cache Session
	redisSource := config.RedisConfig{
		Address: os.Getenv("REDIS_ADDRESS"),
		Port:    os.Getenv("REDIS_PORT"),
	}
	redisStore, _ := redis.NewStore(10, "tcp", fmt.Sprintf("%s:%s", redisSource.Address, redisSource.Port), "", []byte(SECRET))

	dbConn := config.NewDBConnection(DB_SOURCE)
	authRepos := repository.NewAuthRepository(dbConn)
	authController := controller.NewAuthController(authRepos)

	r.LoadHTMLGlob("views/*")

	authController.SetupSession(redisStore, r)
	authController.Routes(r)

	r.Run(":3000")
}
