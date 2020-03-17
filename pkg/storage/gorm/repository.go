package gorm

import (
	"github.com/ajordi/todo/pkg/adding"
	"github.com/jinzhu/gorm"
)

type GormTaskRepository struct {
	DB *gorm.DB
}

func ProvideTaskRepository(DB *gorm.DB) adding.TaskRepository {
	repo := &GormTaskRepository{}
	repo.DB = DB

	return repo
}

func (t *GormTaskRepository) FindAll() []adding.Task {
	var tasks []adding.Task
	t.DB.Find(&tasks)

	return tasks
}

func (t *GormTaskRepository) FindByID(id uint) adding.Task {
	var task adding.Task
	t.DB.First(&task, id)

	return task
}

func (t *GormTaskRepository) Save(task adding.Task) adding.Task {
	t.DB.Save(&task)

	return task
}

func (t *GormTaskRepository) Delete(task adding.Task) {
	t.DB.Delete(&task)
}
