package api

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ajordi/todo/pkg/adding"
	"github.com/ajordi/todo/pkg/authenticating"
	"github.com/ajordi/todo/pkg/deleting"
	"github.com/ajordi/todo/pkg/listing"
	"github.com/google/wire"
)

type Api struct {
	Adding adding.TaskAPI
	AuthMiddleware *jwt.GinJWTMiddleware
	Authenticating authenticating.AuthenticateAPI
	Deleting deleting.TaskAPI
	Listing listing.TaskAPI
}

var ProviderTaskAPI = wire.NewSet(New)

func New(as adding.TaskService, aus authenticating.AuthenticateService, ls listing.TaskService, ds deleting.TaskService) *Api {
	a := &Api{
		Adding :  adding.TaskAPI{TaskService: as},
		AuthMiddleware: authenticating.JWTMiddleware(aus),
		Authenticating: authenticating.AuthenticateAPI{AuthenticateService: aus},
		Deleting : deleting.TaskAPI{TaskService: ds},
		Listing : listing.TaskAPI{TaskService: ls},
	}

	return a
}
