package model

type Product struct {
	ID       int64
	Category Category
	Name     string
	Price    uint64
	ImageUrl string
}

type Category struct {
	ID   int64
	Name string
}
