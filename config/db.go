package config

import (
	"log"
	"notes-v1/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConn(source string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(source), &gorm.Config{})

  if err != nil {
    log.Fatal("Failed to connect database", err)
    return nil
  }
  
  db.AutoMigrate(&model.Note{})
  return db
}
