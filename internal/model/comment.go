package model

import "time"

type Comment struct {
	ID        int64
	IDArticle int64
	Comments  string
	PublishedDate time.Time
}

