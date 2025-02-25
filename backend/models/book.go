package models

type Book struct {
	ID       string   `json:"id,omitempty"`
	Name     string   `json:"name" validate:"required"`
	Subject  string   `json:"subject" validate:"required"`
	Code     string   `json:"code" validate:"required"`
	Slots    []string `json:"slots" validate:"required"`
	OwnerID  string   `json:"ownerId"`
	ImageURL string   `json:"imageUrl,omitempty"`
}
