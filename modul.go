package main

import "time"

// Person ...
type Person struct {
	Firstname string `json:"firstname" binding:"required" `
	Lastname  string `json:"lastname" binding:"required"`
}

// Content ...
type Content struct {
	Title string `json:"title" binding:"required" `
	Body  string `json:"body" binding:"required"`
}

// Article ...
type Article struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	Author    Person     `json:"author" binding:"required" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
