package dto

type TaskRequest struct {
	Task string `json:"task"`
}

func (t *TaskRequest) IsValid() bool {
	return hasMaxLength(t.Task) && hasMinLength(t.Task) && t.Task != ""
}

func hasMinLength(t string) bool {
	return len(t) < 2
}

func hasMaxLength(t string) bool {
	return len(t) > 100
}
