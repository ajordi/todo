package dao

import (
	"github.com/ajordi/todo/pkg/adding"
	"github.com/ajordi/todo/pkg/deleting"
	"github.com/ajordi/todo/pkg/listing"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	storage "github.com/ajordi/todo/pkg/storage/gorm"
)

var ProviderDB = wire.NewSet(New, NewAddingTaskRepository, NewListingTaskRepository, NewDeletingTaskRepository)

func New(GORM *gorm.DB) *storage.Storage {
	s := &storage.Storage{}
	s.GORM = GORM

	return s
}

func NewAddingTaskRepository(s *storage.Storage) adding.TaskRepository {
	return s
}

func NewListingTaskRepository(s *storage.Storage) listing.TaskRepository {
	return s
}

func NewDeletingTaskRepository(s *storage.Storage) deleting.TaskRepository {
	return s
}
