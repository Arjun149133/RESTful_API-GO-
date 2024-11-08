package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content  string `json: "comment"`
	Likes    int32  `json: "likes"`
	PostID   string `json:"post_id"`
	AuthorID string `json:"author_id"`
	Author   Author `gorm: "foreignKey: AuthorID" json: "author"`
	Post     Post   `gorm: "foreignKey: PostID" json: "post"`
	// ParentID *uint  `json: "parent_id"`
	// Replies   []Comment `gorm: "foreignKey: ParentID" json: "replies"`
}

func MigrateComment(db *gorm.DB) {
	db.AutoMigrate(&Comment{})
}
