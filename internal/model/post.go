package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Upvotes int32  `json:"upvotes"`
	// AuthorID  uint      `json: "author_id"`
	// Author    Author    `gorm: "foreignKey: AuthorID" json: "author"`
	// Comments  []Comment `gorm: "embedded" json: "comments"`
}

func MigratePost(db *gorm.DB) {
	db.AutoMigrate(&Post{})
}
