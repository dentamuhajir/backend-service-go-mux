package model

import "time"

type Article struct {
	ID           int64
	Title        string
	Author       string
	Category     string
	ContentBody  string
	Photo        string
	Photographer string
	//IsPublished   bool
	PublishedDate time.Time
	// CreatedAt     time.Time
	// UpdatedAt     time.Time
}

type Headline struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Photo    string `json:"photo"`
}

type ListArticle struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Photo    string `json:"photo"`
}

type PostArticle struct {
	Title         string `json:"title"`
	Author        string
	Category      string
	ContentBody   string `json:"content_body"`
	Photo         string
	Photographer  string
	IsPublished   bool
	PublishedDate time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type StoreArticle struct {
	Title        string
	Author       string
	Category     string
	ContentBody  string
	Photo        string
	Photographer string
}
