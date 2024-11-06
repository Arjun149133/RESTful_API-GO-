package postgres

import (
	"example/restapi/internal/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func (c *CommentRepository) Create(comment *model.Comment) error {
	return c.DB.Create(comment).Error
}

func (c *CommentRepository) FindAll(postId string) ([]model.Comment, error) {
	var comments []model.Comment
	err := c.DB.Preload("Author").Preload("Post").Find(&comments, "post_id = ?", postId).Error
	return comments, err
}

func (c *CommentRepository) FindById(id string) (model.Comment, error) {
	var comment model.Comment
	err := c.DB.Preload("Author").Preload("Post").First(&comment, id).Error
	return comment, err
}

func (c *CommentRepository) Update(comment *model.Comment, postId string) error {
	return c.DB.Save(comment).Error
}

func (c *CommentRepository) Delete(id string) error {
	return c.DB.Delete(&model.Comment{}, id).Error
}
