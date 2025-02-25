package models

type Book struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Code    string `json:"code"`
	Slots   int    `json:"slots"`
	OwnerID string `json:"ownerId"`
}
