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

func (p *PostRepo) FindById(id uint) (model.Post, error) {
	var post model.Post
	res := p.DB.First(&post, "id = ?", id)

	return post, res.Error
}

func (p *PostRepo) Update(post *model.Post) error {
	var oldPost model.Post
	p.DB.First(&oldPost, "id = ?", post.ID)
	return p.DB.UpdateColumns(post).Error
}

func (p *PostRepo) Delete(id uint) error {
	return p.DB.Delete(&model.Post{}, id).Error
}
