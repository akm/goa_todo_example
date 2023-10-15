package todo

import "time"

type Todo struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	State     State     `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
