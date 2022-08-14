package routes

import (
	"github.com/glugate/uno/apps/crm/controllers"
	uno "github.com/glugate/uno/pkg/uno"
	"github.com/glugate/uno/pkg/uno/server"
)

func UsersRoutes() []*server.Route {
	var routes = []*server.Route{
		uno.Get("/users/([^/]+)", controllers.UsersFind),
		uno.Get("/users/?", controllers.UsersList),
	}
	return routes
}
