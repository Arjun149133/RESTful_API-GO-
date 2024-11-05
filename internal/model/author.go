package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name     string    `json: "name" gorm:"not null"`
	Email    string    `gorm:"unique;not null" json: "email"`
	Posts    []Post    `json: "posts"`
	Comments []Comment `json: "comments"`
}

func MigrateAuthor(db *gorm.DB) {
	db.AutoMigrate(&Author{})
}
