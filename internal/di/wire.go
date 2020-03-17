//+build wireinject

package di

import (
	"github.com/ajordi/todo/pkg/adding"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitTaskAPI(db *gorm.DB) adding.TaskAPI {
	wire.Build(ProviderDB, adding.ProvideTaskService, adding.ProvideTaskAPI)
	return adding.TaskAPI{}
}
