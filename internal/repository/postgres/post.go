package postgres

import (
	"example/restapi/internal/model"

	"gorm.io/gorm"
)

type PostRepo struct {
	DB *gorm.DB
}

func (p *PostRepo) Create(post *model.Post) error {
	return p.DB.Create(post).Error
}

func (p *PostRepo) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := p.DB.Find(&posts).Error
	return posts, err
}
