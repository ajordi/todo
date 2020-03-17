package deleting

type TaskRepository interface {
	Delete(task Task)
}
