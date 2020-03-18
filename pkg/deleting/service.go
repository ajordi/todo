package deleting

type TaskService interface {
	Delete(id uint)
}

type taskService struct {
	TaskRepository TaskRepository
}

func ProvideTaskService(t TaskRepository) TaskService {
	return &taskService{t}
}

func (t *taskService) Delete(id uint) {
	t.TaskRepository.Delete(id)
}

