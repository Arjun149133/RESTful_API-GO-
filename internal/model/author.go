package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	ID       string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `json:"password" gorm:"not null"`
	Posts    []Post    `gorm:"foreignKey:AuthorID" json:"posts"`
	Comments []Comment `gorm:"foreignKey:AuthorID" json:"comments"`
}

func MigrateAuthor(db *gorm.DB) {
	db.AutoMigrate(&Author{})
}

// Helper function to generate a new UUID string
func NewUUID() string {
	return uuid.New().String()
}
