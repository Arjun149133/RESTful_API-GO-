package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	// ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Upvotes  int32     `json:"upvotes"`
	AuthorID uint      `json: "author_id"`
	Author   Author    `gorm: "foreignKey: AuthorID" json: "author"`
	Comments []Comment `gorm: "embedded" json: "comments"`
}

func MigratePost(db *gorm.DB) {
	db.AutoMigrate(&Post{})
}
