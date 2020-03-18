package listing

func ToTask(taskDTO TaskDTO) Task {
	return Task{Name: taskDTO.Name}
}

func ToTaskDTO(task Task) TaskDTO {
	return TaskDTO{ID: task.ID, Name: task.Name}
}

func ToTaskDTOs(tasks []Task) []TaskDTO {
	taskdtos := make([]TaskDTO, len(tasks))

	for i, itm := range tasks {
		taskdtos[i] = ToTaskDTO(itm)
	}

	return taskdtos
}
