package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server is running")
  r := gin.Default()
  r.Run(":8080")
}
