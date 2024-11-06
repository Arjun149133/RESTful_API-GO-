package service

import (
	"errors"
	"example/restapi/internal/model"
	"example/restapi/internal/repository"
)

type AuthorService struct {
	Repo repository.AuthorRepository
}

func (s *AuthorService) CreateAuthor(author *model.Author) error {
	if author.Email == "" {
		return errors.New("email cannot be empty")
	}
	if author.Name == "" {
		return errors.New("name cannot be empty")
	}
	author.ID = model.NewUUID()

	return s.Repo.Create(author)
}
func (s *AuthorService) GetAllAuthors() ([]model.Author, error) {
	return s.Repo.FindAll()
}
func (s *AuthorService) GetAuthorById(id string) (*model.Author, error) {
	return s.Repo.FindById(id)
}
func (s *AuthorService) UpdateAuthor(author *model.Author) error {
	if author.Email == "" {
		return errors.New("email cannot be empty")
	}
	if author.Name == "" {
		return errors.New("name cannot be empty")
	}
	return s.Repo.Update(author)
}
func (s *AuthorService) DeleteAuthor(id string) error {
	return s.Repo.Delete(id)
}
