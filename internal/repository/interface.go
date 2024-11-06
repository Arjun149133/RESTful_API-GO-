package repository

import "example/restapi/internal/model"

type AuthorRepository interface {
	Create(author *model.Author) error
	FindAll() ([]model.Author, error)
	FindById(id string) (*model.Author, error)
	Update(author *model.Author) error
	Delete(id string) error
}

type PostRepository interface {
	Create(post *model.Post) error
	FindAll() ([]model.Post, error)
	FindById(id string) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id string) error
}

type CommentRepository interface {
	Create(comment *model.Comment) error
	FindAll(id string) ([]model.Comment, error)
	FindById(postId string) (model.Comment, error)
	Update(comment *model.Comment, postId string) error
	Delete(id string) error
}
