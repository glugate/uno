package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/glugate/uno/pkg/uno/log"
)

// Server is internal API server for handling requests
type Server struct {
	stdServer *http.Server
	Router    *Router
}

// Creates new server instance on a given http address
func NewServer() *Server {
	address, _ := os.LookupEnv("ADDRESS")
	port, _ := os.LookupEnv("PORT")
	srv := &Server{
		stdServer: &http.Server{
			Addr: fmt.Sprintf("%s:%s", address, port),
		},
		Router: NewRouter(),
	}

	// setup server handler
	srv.stdServer.Handler = srv.Router

	return srv
}

// RegisterRoutes registeres all or partial routes
// that are passed and kept in Router
func (s *Server) RegisterRoutes(routes []*Route) {
	s.Router.RegisterRoutes(routes)
}

// Run starts the http server on specified port in .env file
func (o *Server) Run() {
	address, exists := os.LookupEnv("ADDRESS")
	if !exists {
		address = ServerDefaultAddress
	}
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = ServerDefaultPort
	}
	url := fmt.Sprintf("%s:%s", address, port)
	log.Default().Verbose("There are %d routes defined.", len(o.Router.Routes))
	log.DefaultLogFactory().Default().Info("Server running at: %S", url)
	if err := o.stdServer.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
}
