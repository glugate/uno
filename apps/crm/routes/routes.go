package routes

import "github.com/glugate/uno/pkg/uno/server"

func All() []*server.Route {
	return MenusRoutes()
}
