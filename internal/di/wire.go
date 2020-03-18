//+build wireinject

package di

import (
	"github.com/ajordi/todo/internal/api"
	"github.com/ajordi/todo/internal/dao"
	"github.com/ajordi/todo/internal/service"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitTaskAPI(db *gorm.DB) *api.Api {
	wire.Build(dao.ProviderDB, service.ProviderService, api.ProviderTaskAPI)
	return &api.Api{}
}
