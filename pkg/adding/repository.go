package adding

type TaskRepository interface {
	FindAll() []Task
	FindByID(id uint) Task
	Save(task Task) Task
	Delete(task Task)
}




