package config

import (
	"fmt"

	"notes-v1/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConn(source string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(source), &gorm.Config{})

  if err != nil {
    fmt.Println("Failed to connect database", err)
    return nil
  }
  
  err = db.AutoMigrate(&model.Note{})

  if err != nil {
    fmt.Println("Failed to migrate database")
    return nil
  }
  return db
}
