package api

import (
	"github.com/ajordi/todo/pkg/adding"
	"github.com/ajordi/todo/pkg/deleting"
	"github.com/ajordi/todo/pkg/listing"
	"github.com/google/wire"
)

type Api struct {
	Adding adding.TaskAPI
	Listing listing.TaskAPI
	Deleting deleting.TaskAPI
}

var ProviderTaskAPI = wire.NewSet(New)

func New(as adding.TaskService, ls listing.TaskService, ds deleting.TaskService) *Api {
	a := &Api{
		Adding :  adding.TaskAPI{TaskService: as},
		Listing : listing.TaskAPI{TaskService: ls},
		Deleting : deleting.TaskAPI{TaskService: ds},
	}

	return a
}
