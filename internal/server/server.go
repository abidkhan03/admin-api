package server

import (
	"fmt"
	"github.com/spongeling/admin-api/internal/api"
	"log"
	"net/http"
)

type Server struct {
	port int
	apis []api.API
}

func New(port int, apis ...api.API) *Server {
	return &Server{port: port, apis: apis}
}

// Start starts the api server
func (srv *Server) Start() error {
	r := srv.buildRouter()

	httpPort := fmt.Sprintf(":%d", srv.port)
	log.Printf("Starting server on %v\n", httpPort)

	return http.ListenAndServe(httpPort, r)
}
