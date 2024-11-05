package repository

import "example/restapi/internal/model"

type AuthorRepository interface {
	Create(author *model.Author) error
	FindAll() ([]model.Author, error)
	FindById(id uint) (model.Author, error)
	Update(author *model.Author) error
	Delete(id uint) error
}

type PostRepository interface {
	Create(post *model.Post) error
	FindAll() ([]model.Post, error)
	FindById(id uint) (model.Post, error)
	Update(post *model.Post) error
	Delete(id uint) error
}
