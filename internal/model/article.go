package model

import "time"

type Article struct {
	ID            int64
	Title         string
	Author        string
	Category      string
	ContentBody   string
	Photo         string
	Photographer  string
	//IsPublished   bool
	PublishedDate time.Time
	// CreatedAt     time.Time
	// UpdatedAt     time.Time
}
