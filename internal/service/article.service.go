package service

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
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

func (s *ArticleService) GetListArticle() ([]model.Article, error) {
	return s.articleRepository.GetArticleList()
}

func (s *ArticleService) GetListArticleGroupByCategory() (map[string][]model.Article, error) {
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

func (s *ArticleService) GetHeadlineArticle() ([]model.Article, error) {
	return s.articleRepository.GetHeadlineArticle()
}

func (s *ArticleService) GenerateDummyData() (message string) {
	store := model.StoreArticle{}

	store.Title = gofakeit.Book().Title
	store.Author = gofakeit.Name()
	store.Category = gofakeit.BookGenre()
	store.ContentBody = gofakeit.Word()
	store.Photo = ""
	store.Photographer = gofakeit.BeerName()

	fmt.Println(store)

	message = s.articleRepository.SaveArticle(store)

	return message

}
