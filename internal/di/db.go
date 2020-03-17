package di

import (
	"github.com/ajordi/todo/pkg/adding"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	storage "github.com/ajordi/todo/pkg/storage/gorm"
)

var ProviderDB = wire.NewSet(New, NewAddingTaskRepository)

func New(GORM *gorm.DB) *storage.Storage {
	s := &storage.Storage{}
	s.GORM = GORM

	return s
}

func NewAddingTaskRepository(s *storage.Storage) adding.TaskRepository {
	return s
}
