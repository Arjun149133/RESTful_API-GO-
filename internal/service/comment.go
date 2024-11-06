package service

import (
	"errors"
	"example/restapi/internal/model"
	"example/restapi/internal/repository"
)

type CommentService struct {
	Repo repository.CommentRepository
}

func (s *CommentService) CreateComment(comment *model.Comment) error {
	if comment.Content == "" {
		return errors.New("content cannot be empty")
	}
	if comment.PostID == "" {
		return errors.New("post_id cannot be empty")
	}
	comment.ID = model.NewUUID()
	return s.Repo.Create(comment)
}

func (s *CommentService) GetAllComments(postId string) ([]model.Comment, error) {
	return s.Repo.FindAll(postId)
}

func (s *CommentService) GetCommentById(id string) (model.Comment, error) {
	return s.Repo.FindById(id)
}

func (s *CommentService) UpdateComment(comment *model.Comment, postId string) error {
	if comment.Content == "" {
		return errors.New("content cannot be empty")
	}
	if comment.PostID != postId {
		return errors.New("post_id cannot be changed")
	}
	return s.Repo.Update(comment, postId)
}

func (s *CommentService) DeleteComment(id string) error {
	return s.Repo.Delete(id)
}
