package uno

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/glugate/uno/pkg/uno/db"
	"github.com/glugate/uno/pkg/uno/server"
)

// Uno
type Uno struct {
	DB     *db.DB
	Server *server.Server
}

// NewUno
func NewUno() *Uno {
	return &Uno{
		DB:     db.NewDB(),
		Server: server.NewServer(),
	}
}

// Get creates new route with GET method and passed pattern and handler
func Get(pattern string, handler http.HandlerFunc) *server.Route {
	return &server.Route{Method: "GET", Regex: regexp.MustCompile("^" + pattern + "$"), Handler: handler}
}

// Post creates new route with POST method and passed pattern and handler
func Post(pattern string, handler http.HandlerFunc) *server.Route {
	return &server.Route{Method: "POST", Regex: regexp.MustCompile("^" + pattern + "$"), Handler: handler}
}

func (o *Uno) Metrics() {
	fmt.Printf("Metrics: %s\n", "TODO")
}
