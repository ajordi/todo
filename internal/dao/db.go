package dao

import (
	"github.com/ajordi/todo/pkg/adding"
	"github.com/ajordi/todo/pkg/authenticating"
	"github.com/ajordi/todo/pkg/deleting"
	"github.com/ajordi/todo/pkg/listing"
	"github.com/ajordi/todo/pkg/storage/gorm/task"
	"github.com/ajordi/todo/pkg/storage/gorm/user"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var ProviderDB = wire.NewSet(NewTaskStorage, NewUserStorage, NewAddingTaskRepository, NewAuthenticatingUserRepository, NewListingTaskRepository, NewDeletingTaskRepository)

func NewTaskStorage(GORM *gorm.DB) *task.Storage {
	s := &task.Storage{}
	us := &user.Storage{}
	s.GORM = GORM
	us.GORM = GORM

	return s
}

func NewUserStorage(GORM *gorm.DB) *user.Storage {

	s := &user.Storage{}
	s.GORM = GORM

	return s
}


func NewAddingTaskRepository(s *task.Storage) adding.TaskRepository {
	return s
}

func NewAuthenticatingUserRepository(s *user.Storage) authenticating.UserRepository{
	return s
}

func NewListingTaskRepository(s *task.Storage) listing.TaskRepository {
	return s
}

func NewDeletingTaskRepository(s *task.Storage) deleting.TaskRepository {
	return s
}
