package events

import (
	"time"
)

type Message interface {
	Type() string
}

type CreatedFeedMessage struct {
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (m CreatedFeedMessage) Type() string {
	return "created_feed"
}
