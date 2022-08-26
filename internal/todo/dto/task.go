package dto

type TaskRequest struct {
	Task string `json:"task"`
}

const (
	maxLength = 100
	minLength = 2
)

func (t *TaskRequest) IsValid() bool {
	return hasMaxLength(t.Task) && hasMinLength(t.Task) && t.Task != ""
}

func hasMinLength(t string) bool {
	return len(t) > minLength
}

func hasMaxLength(t string) bool {
	return len(t) < maxLength
}
