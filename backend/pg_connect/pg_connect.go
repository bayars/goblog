package pg_connect

import (
	"fmt"
	"os"

	"goblog/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

// GetData retrieves an article by its index number.
func GetData(index uint) (*models.Article, error) {
	var article models.Article

	// Query the database for the article with the specified index
	result := Database.First(&article, index)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}

// Get data function should search database with gorm and get the article with data index number
