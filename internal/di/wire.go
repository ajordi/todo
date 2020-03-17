//+build wireinject

package di

import (
	"github.com/ajordi/todo/pkg/adding"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	storage "github.com/ajordi/todo/pkg/storage/gorm"
)

func InitTaskAPI(db *gorm.DB) adding.TaskAPI {
	wire.Build(storage.ProvideTaskRepository, adding.ProvideTaskService, adding.ProvideTaskAPI)
	return adding.TaskAPI{}
}
