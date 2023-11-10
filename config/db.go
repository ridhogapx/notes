package config

import (
	"context"
	"fmt"
	"notes-v1/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConn(source string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(source), &gorm.Config{})

  // Setup context time out
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()

  db.WithContext(ctx)

  if err != nil {
    fmt.Println("Failed to connect database", err)
    return nil
  }
 
  db.AutoMigrate(&model.Note{})

  return db
}
