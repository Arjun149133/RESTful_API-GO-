package service

import (
	"example/restapi/internal/model"
	"example/restapi/internal/repository"
)

type PostService struct {
	Repo repository.PostRepository
}

func (s *PostService) CreatePost(post *model.Post) error {
	return s.Repo.Create(post)
}

func (s *PostService) GetAllPosts() ([]model.Post, error) {
	return s.Repo.FindAll()
}

func (s *PostService) GetPostById(id uint) (model.Post, error) {
	return s.Repo.FindById(id)
}

func (s *PostService) UpdatePost(post *model.Post) error {
	return s.Repo.Update(post)
}

func (s *PostService) DeletePost(id uint) error {
	return s.Repo.Delete(id)
}
