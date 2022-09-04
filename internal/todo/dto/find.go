package dto

import (
	"net/url"
	"time"

	"github.com/CaioAureliano/go-do/internal/todo/model"
)

type FindResponse struct {
	Todos []*model.Todo `json:"todos"`
	Count int           `json:"count"`
}

type FilterRequest struct {
	Task           string     `json:"task"`
	Status         *bool      `json:"status"`
	DateStartRange *time.Time `json:"date_start"`
	DateEndRange   *time.Time `json:"date_end"`

	SortByField string `json:"sort_by"`
}

var (
	newBool = func(b bool) *bool { return &b }
)

const defaultTimeFormat = "01/02/2006"

func (f *FilterRequest) Mount(q url.Values) {
	if q.Has("task") {
		f.Task = q.Get("task")
	}

	if q.Has("status") {
		f.Status = newBool(false)
		if q.Get("status") == "true" {
			f.Status = newBool(true)
		}
	}

	if q.Has("date_start") {
		t, _ := time.Parse(defaultTimeFormat, q.Get("date_start"))
		f.DateStartRange = &t
	}

	if q.Has("date_end") {
		t, _ := time.Parse(defaultTimeFormat, q.Get("date_end"))
		f.DateEndRange = &t
	}
}
