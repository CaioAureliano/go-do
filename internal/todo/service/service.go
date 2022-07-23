package service

type TodoService interface {
	Create(task string) error
}

type todoService struct {
}

func New() TodoService {
	return todoService{}
}

func (t todoService) Create(task string) error {
	return nil
}
