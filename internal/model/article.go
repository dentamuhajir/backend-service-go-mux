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
	PublishedDate time.Time
}
