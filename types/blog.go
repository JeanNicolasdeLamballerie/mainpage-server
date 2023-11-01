package typing

import "time"

type Post struct {
	Unique_idx  string
	Name        string
	Date        time.Time
	Description string
	Content     string
}

type BlogCategory struct {
	Unique_idx string
	Name       string
	Latest     struct {
		Unique_idx  string
		Name        string
		Date        time.Time
		Description string
	}
}

type BlogData struct {
	mock       bool
	categories []BlogCategory
}
type BlogDatabase struct {
	GetCategories func() []BlogCategory
	GetAllPosts   func() []Post
	GetPost       func(idx string) []Post
	GetCategory   func(idx string) []BlogCategory
}

type DB struct {
	Blog BlogDatabase
}
