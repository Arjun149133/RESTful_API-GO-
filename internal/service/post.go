package service

import (
	"errors"
	"example/restapi/internal/model"
	"example/restapi/internal/repository"
)

type PostService struct {
	Repo repository.PostRepository
}

func (s *PostService) CreatePost(post *model.Post) error {
	if post.Title == "" {
		return errors.New("title cannot be empty")
	}
	if post.Content == "" {
		return errors.New("content cannot be empty")
	}
	return s.Repo.Create(post)
}

func (s *PostService) GetAllPosts() ([]model.Post, error) {
	return s.Repo.FindAll()
}

func (s *PostService) GetPostById(id uint) (model.Post, error) {
	return s.Repo.FindById(id)
}

func (s *PostService) UpdatePost(post *model.Post) error {
	if post.Title == "" {
		return errors.New("title cannot be empty")
	}
	if post.Content == "" {
		return errors.New("content cannot be empty")
	}
	return s.Repo.Update(post)
}

func (s *PostService) DeletePost(id uint) error {
	return s.Repo.Delete(id)
}
