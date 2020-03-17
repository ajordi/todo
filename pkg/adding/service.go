package adding

type TaskService interface {
	Save(task Task) Task
}

type taskService struct {
	TaskRepository TaskRepository
}

func ProvideTaskService(t TaskRepository) TaskService {
	return &taskService{t}
}

func (t *taskService) Save(task Task) Task {
	t.TaskRepository.Save(task)

	return task
}
