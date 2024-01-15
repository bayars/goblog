package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100)" json:"title"`
	Content   string `gorm:"type:text" json:"content"`
	Tags      string `gorm:"type:varchar(100)" json:"tags"`
	Timestamp string `gorm:"type:timestamp" json:"timestamp"`
}

func UseLibFunction() {
	fmt.Println("Using a function from the lib package")
}
