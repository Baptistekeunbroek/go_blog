package models

import "time"

// Post struct to represent a blog post
type Post struct {
	ID      string
	Title   string
	Content string
	Author  string
	Created time.Time
}
