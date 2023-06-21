package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spongeling/admin-api/internal/service"
)

type Ping struct {
	svc *service.Service
}

func NewPing(svc *service.Service) *Ping {
	return &Ping{svc: svc}
}

func (api *Ping) Routes(r chi.Router) {
	r.Get("/hello", api.Hello)
}

func (api *Ping) Hello(w http.ResponseWriter, _ *http.Request) {
	respondOk(w, "everything is working fine..!")
}
