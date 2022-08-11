package dto

import (
	"time"

	"github.com/CaioAureliano/go-do/internal/todo/model"
)

type FindResponse struct {
	Todos []*model.Todo `json:"todos"`
	Count int64         `json:"count"`
}

type FilterRequest struct {
	Task           string    `json:"task"`
	Status         *bool     `json:"status"`
	DateStartRange time.Time `json:"date_start"`
	DateEndRange   time.Time `json:"date_end"`

	SortByField string `json:"sort_by"`
}
