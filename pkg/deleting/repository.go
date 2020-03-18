package deleting

type TaskRepository interface {
	Delete(id uint)
}
