package models

// Chat model
type Chat struct {
	ID        string        `json:"id"`
	BookID    string        `json:"book_id"`
	Requester string        `json:"requester"`
	Owner     string        `json:"owner"`
	Messages  []ChatMessage `json:"messages"`
}

// ChatMessage model
type ChatMessage struct {
	Sender    string `json:"sender"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}
