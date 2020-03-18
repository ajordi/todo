package service

import (
	"github.com/ajordi/todo/pkg/adding"
	"github.com/ajordi/todo/pkg/deleting"
	"github.com/ajordi/todo/pkg/listing"
	"github.com/google/wire"
)

var ProviderService = wire.NewSet(adding.ProvideTaskService, listing.ProvideTaskService, deleting.ProvideTaskService)
