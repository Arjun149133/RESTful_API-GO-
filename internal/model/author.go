package model

import "gorm.io/gorm"

type Author struct {
	ID       uint      `json: "id"`
	Name     string    `json: "name"`
	Email    string    `json: "email" gorm:"unique"`
	Posts    []Post    `json: "posts"`
	Comments []Comment `json: "comments"`
}

func MigrateAuthor(db *gorm.DB) {
	db.AutoMigrate(&Author{})
}
