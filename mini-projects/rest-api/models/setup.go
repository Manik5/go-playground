package models

import (
  "github.com/jinzhu/gorm"
  // because yes
  _"github.com/jinzhu/gorm/dialects/sqlite"
)

// DB var
var DB *gorm.DB

// ConnectDatabase func
func ConnectDatabase() {
    database, err := gorm.Open("sqlite", "test.db")

    if err != nil {
      panic("Failed to connect to database!")
  }

  database.AutoMigrate(&Book{})

  DB = database
}

