package models

import "time"

type Request struct {
	ID        string    `json:"id,omitempty"`
	BookID    string    `json:"book_id" validate:"required"`
	Requester string    `json:"requester" validate:"required"` // User requesting the book
	Owner     string    `json:"owner" validate:"required"`     // Book owner
	Status    string    `json:"status"`                        // pending, accepted, rejected
	CreatedAt time.Time `json:"created_at"`
}
