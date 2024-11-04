package repository

import "example/restapi/internal/model"

type PostRepository interface {
	Create(post *model.Post) error
	FindAll() ([]model.Post, error)
	FindById(id uint) (model.Post, error)
	Update(post *model.Post) error
	Delete(id uint) error
}
