package entities

import "time"

type Timed struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
