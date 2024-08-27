package article

import (
	"log"

	"github.com/dentamuhajir/backend-service-go-mysql/models"
	"github.com/dentamuhajir/backend-service-go-mysql/utils"
)

type ArticleService struct {
	articleRepository ArticleRepository
}

func NewArticleService(articleRepository ArticleRepository) *ArticleService {
	return &ArticleService{articleRepository}
}

func (s ArticleService) postArticle(requestBody models.Article) string {
	message := s.articleRepository.insertArticle(requestBody)

	log.Println(message)
	// if err != nil {
	// 	return nil
	// }
	return message
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
