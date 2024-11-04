package repository

import "example/restapi/internal/model"

type PostRepository interface {
	Create(post *model.Post) error
	FindAll() ([]model.Post, error)
}
