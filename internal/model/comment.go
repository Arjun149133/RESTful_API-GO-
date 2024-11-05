package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content  string `json: "comment"`
	Likes    int32  `json: "likes"`
	PostID   uint   `json:"post_id"`
	AuthorID uint   `json:"author_id"`
	ParentID *uint  `json: "parent_id"`
	Author   Author `gorm: "foreignKey: AuthorID" json: "author"`
	Post     Post   `gorm: "foreignKey: PostID" json: "post"`
	// Replies   []Comment `gorm: "foreignKey: ParentID" json: "replies"`
}

func MigrateComment(db *gorm.DB) {
	db.AutoMigrate(&Comment{})
}
