package dto

import "time"

type CommentDTO struct {
	ID        uint      `json:"id"`
	Author    string    `json:"author" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	CreatedAt time.Time `json:"created-at"`
}
