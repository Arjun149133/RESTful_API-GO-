package model

import (
	"gorm.io/gorm"
)

type Post struct {
	ID       string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Upvotes  int32     `json:"upvotes"`
	AuthorID string    `json: "author_id"`
	Author   Author    `gorm: "foreignKey: AuthorID" json: "author"`
	Comments []Comment `gorm: "foreignKey: PostID" json: "comments"`
}

func MigratePost(db *gorm.DB) {
	db.AutoMigrate(&Post{})
}
