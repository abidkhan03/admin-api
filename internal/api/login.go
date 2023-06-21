package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spongeling/admin-api/internal/auth"
	"github.com/spongeling/admin-api/internal/response"
	"github.com/spongeling/admin-api/internal/service"
)

type Login struct {
	svc *service.Service
}

func NewLogin(svc *service.Service) *Login {
	return &Login{svc: svc}
}

func (api *Login) Routes(r chi.Router) {
	r.Post("/login", auth.Authenticator(api.Login))
}

// Login is a handler for POST /login
func (api *Login) Login(w http.ResponseWriter, r *http.Request) {
	// respond
	respondOk(w, response.Response{
		Status:  http.StatusOK,
		Message: "successfully logged in",
	})
}
