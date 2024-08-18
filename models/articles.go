package models

type Article struct {
	ID         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Body       string `json:"body" db:"body"`
	Published  bool   `json:"published" db:"published"`
	CategoryID int    `json:"category_id" db:"category_id"`
}
