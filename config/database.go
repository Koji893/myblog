/ config/database.go
package config

import (
	"log"
	"my-go-blog/models" // Make sure this import path is correct!
	"gorm.io/driver/sqlite" // Or postgres for production
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{}) // Or your PostgreSQL connection
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// This line tells GORM to create or update the 'posts' table based on your Post struct
	err = database.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	DB = database
}
