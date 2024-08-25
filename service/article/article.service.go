package article

import (
	"github.com/dentamuhajir/backend-service-go-mysql/models"
	"github.com/dentamuhajir/backend-service-go-mysql/utils"
)

type ArticleService struct {
	articleRepository ArticleRepository
}

func NewArticleService(articleRepository ArticleRepository) *ArticleService {
	return &ArticleService{articleRepository}
}

func (s ArticleService) getArticle() []models.Article {
	articles, err := s.articleRepository.getArticle()

	if err != nil {
		return []models.Article{}
	}
	for i := range articles {
		articles[i].Slug = utils.GenerateSlug(articles[i].Title)
	}

	return articles
}
