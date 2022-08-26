package model

import "time"

type Todo struct {
	ID        string
	Task      string
	Status    bool
	CreatedAt time.Time
}
