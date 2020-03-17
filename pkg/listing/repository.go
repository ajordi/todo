package listing

type TaskRepository interface {
	FindAll() []Task
	FindByID(id uint) Task
}
