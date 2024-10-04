package models

import "time"

// Post struct to represent a blog post
type Post struct {
	ID      string
	Title   string
	Content string
	Author  string
	Image   string
	Created time.Time
	UserID  string
}
