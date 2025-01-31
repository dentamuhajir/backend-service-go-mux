package service

import (
	"github.com/dentamuhajir/backend-service-go-mysql/internal/model"
	"github.com/dentamuhajir/backend-service-go-mysql/internal/repository"
)

type CommentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) *CommentService {
	return &CommentService{commentRepository: r}
}

func (s *CommentService) GetCommentByArticleID(articleId int64) ([]model.Comment, error) {
	comments, err := s.commentRepository.GetCommentsByArticleID(articleId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
