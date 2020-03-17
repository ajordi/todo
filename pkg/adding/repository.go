package adding

type TaskRepository interface {
	Save(task Task) Task
}
