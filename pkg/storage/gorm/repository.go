package gorm

import (
	"github.com/ajordi/todo/pkg/adding"
	"github.com/ajordi/todo/pkg/deleting"
	"github.com/ajordi/todo/pkg/listing"
	"github.com/jinzhu/gorm"
)

type Storage struct {
	GORM *gorm.DB
}

func (t *Storage) FindAll() []listing.Task {
	var tasks []listing.Task
	t.GORM.Find(&tasks)

	return tasks
}

func (t *Storage) FindByID(id uint) listing.Task {
	var task listing.Task
	t.GORM.First(&task, id)

	return task
}

func (t *Storage) Save(task adding.Task) adding.Task {
	t.GORM.Save(&task)

	return task
}

func (t *Storage) Delete(task deleting.Task) {
	t.GORM.Delete(&task)
}
