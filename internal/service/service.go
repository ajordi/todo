package service

import (
	"github.com/ajordi/todo/pkg/adding"
	"github.com/ajordi/todo/pkg/authenticating"
	"github.com/ajordi/todo/pkg/deleting"
	"github.com/ajordi/todo/pkg/listing"
	"github.com/google/wire"
)

var ProviderService = wire.NewSet(adding.ProvideTaskService, authenticating.ProviderAuthenticatingService, listing.ProvideTaskService, deleting.ProvideTaskService)
