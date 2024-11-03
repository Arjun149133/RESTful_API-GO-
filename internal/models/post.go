package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Upvotes   int32     `json:"upvotes"`
	AuthorID  uint      `json: "author_id"`
	Author    Author    `gorm: "foreignKey: AuthorID" json: "author"`
	Comments  []Comment `gorm: "embedded" json: "comments"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func MigratePost(db *gorm.DB) {
	db.AutoMigrate(&Post{})
}
