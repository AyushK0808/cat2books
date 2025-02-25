package models

type Request struct {
	ID          string `json:"id"`
	BookID      string `json:"bookId"`
	RequesterID string `json:"requesterId"`
	Status      string `json:"status"`
}
