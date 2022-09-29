package content

import "time"

type Content struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:description`
	CreatedAt   time.Time `json:"created_at"`
}
