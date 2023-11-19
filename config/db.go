package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(source string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(source), &gorm.Config{})

  if err != nil {
    fmt.Println("Failed to connect database:", err)
    return nil
  }

  return db
}
