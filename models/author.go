package models

import "time"

// Author ..
type Author struct {
	ID        string     `json:"id"`
	Firstname string     `json:"firstname" binding:"required" minLength:"4" maxLength:"50" example:"John--example"`
	Lastname  string     `json:"lastname" binding:"required" minLength:"4" maxLength:"50" example:"Does--example"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// CreateAuthorModul ..
type CreateAuthorModul struct {
	Firstname string `json:"firstname" binding:"required" minLength:"4" maxLength:"50" example:"John"`
	Lastname  string `json:"lastname" binding:"required" minLength:"4" maxLength:"50" example:"Does"`
}

// PackedAuthorModel ... bu keremas
type PackedAuthorModel struct {
	ID        string     `json:"id"`
	Firstname string     `json:"firstname" binding:"required" minLength:"4" maxLength:"50" example:"John"`
	Lastname  string     `json:"lastname" binding:"required" minLength:"4" maxLength:"50" example:"Does"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// UpdateAuthorModul ..
type UpdateAuthorModul struct {
	ID        string `json:"id" binding:"required"`
	Firstname string `json:"firstname" binding:"required" minLength:"4" maxLength:"50" example:"John"`
	Lastname  string `json:"lastname" binding:"required" minLength:"4" maxLength:"50" example:"Does"`
}
