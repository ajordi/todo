package adding

func ToTask(taskDTO TaskDTO) Task {
	return Task{Name: taskDTO.Name}
}

func ToTaskDTO(task Task) TaskDTO {
	return TaskDTO{ID: task.ID, Name: task.Name}
}