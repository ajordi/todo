package listing

type TaskService interface {
	FindAll() []Task
	FindByID(id uint) Task
}

type taskService struct {
	TaskRepository TaskRepository
}

func ProvideTaskService(t TaskRepository) TaskService {
	return &taskService{t}
}

func (t *taskService) FindAll() []Task {
	return t.TaskRepository.FindAll()
}

func (t *taskService) FindByID(id uint) Task {
	return t.TaskRepository.FindByID(id)
}
