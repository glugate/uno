package routes

import (
	"github.com/glugate/uno/apps/crm/controllers"
	uno "github.com/glugate/uno/pkg/uno"
	"github.com/glugate/uno/pkg/uno/server"
)

func MenusRoutes() []*server.Route {
	var routes = []*server.Route{
		/*
			uno.Get("/menu/([^/]+)/items/([0-9]+)/update", MenusItemsUpdate),
		*/

		uno.Get("/menus/([^/]+)", controllers.MenusFind),
		uno.Get("/menus", controllers.MenusList),
	}
	return routes
}
