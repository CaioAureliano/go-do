package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string

		gotTask *TaskRequest

		wantValid bool
	}{
		{
			name: "should be true with valid task",

			gotTask: &TaskRequest{
				Task: "GoLang",
			},

			wantValid: true,
		},
		{
			name: "should be true with invalid task greater max range",

			gotTask: &TaskRequest{
				Task: "Gooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo lang",
			},

			wantValid: false,
		},
		{
			name: "should be true with invalid task less min range",

			gotTask: &TaskRequest{
				Task: "Go",
			},

			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := tt.gotTask.IsValid()

			assert.Equal(t, tt.wantValid, isValid)
		})
	}
}
