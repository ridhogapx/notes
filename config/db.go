package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConn(source string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(source), &gorm.Config{})

  if err != nil {
    log.Fatal("Failed to connect database", err)
    return nil
  }

  return db
}
