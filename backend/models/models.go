package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100)" json:"title"`
	Content   string `gorm:"type:text" json:"content"`
	Active    bool   `gorm:"type:boolean" json:"active"`
	Tags      string `gorm:"type:varchar(100)" json:"tags"`
	Timestamp string `gorm:"type:timestamp" json:"timestamp"`
}

func UseLibFunction() {
	fmt.Println("Using a function from the lib package")
}
