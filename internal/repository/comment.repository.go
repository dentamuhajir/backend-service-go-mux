package repository

import (
	"database/sql"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/model"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) GetCommentsByArticleID(articleId int64) ([]model.Comment, error) {
	query := "SELECT * from comments where article_id = ?"
	rows, err := r.db.Query(query, articleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []model.Comment
	for rows.Next() {
		var comment model.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.IDArticle,
			&comment.Comments,
			&comment.PublishedDate,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
