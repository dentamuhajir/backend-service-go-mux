package service

import (
	"github.com/dentamuhajir/backend-service-go-mysql/internal/model"
	"github.com/dentamuhajir/backend-service-go-mysql/internal/repository"
)

type ArticleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(r repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		articleRepository: r,
	}
}

func (s ArticleService) GetListArticle() ([]model.Article, error) {
	return s.articleRepository.GetArticleList()
}

func (s ArticleService) GetListArticleGroupByCategory() (map[string][]model.Article, error) {
	articles, err := s.articleRepository.GetArticleList()
	if err != nil {
		return nil, err
	}
	collections := make(map[string][]model.Article)
	for _, article := range articles {
		collections[article.Category] = append(collections[article.Category], article)
	}
	return collections, nil
}

func (s ArticleService) GetHeadlineArticle() ([]model.Article, error) {
	return s.articleRepository.GetHeadlineArticle()
}
