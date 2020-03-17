package adding

type TaskService interface {
	FindAll() []Task
	FindByID(id uint) Task
	Save(task Task) Task
	Delete(task Task)
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

func (t *taskService) Save(task Task) Task {
	t.TaskRepository.Save(task)

	return task
}

func (t *taskService) Delete(task Task) {
	t.TaskRepository.Delete(task)
}
