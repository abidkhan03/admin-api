package api

import (
	"github.com/go-chi/chi"
	"github.com/spongeling/admin-api/internal/auth"
	"github.com/spongeling/admin-api/internal/request"
	"github.com/spongeling/admin-api/internal/service"
	"log"
	"net/http"
)

type Pattern struct {
	svc *service.Service
}

func NewPattern(svc *service.Service) *Pattern {
	return &Pattern{svc: svc}
}

func (api *Pattern) Routes(r chi.Router) {
	r.Route("/pattern", func(r chi.Router) {
		r.Post("/generate", auth.Authenticator(api.GetPhrasePattern))
		r.Post("/examples", auth.Authenticator(api.GetPatternExamples))
	})
}

// GetPhrasePattern is handler for route POST /pattern/generate
func (api *Pattern) GetPhrasePattern(w http.ResponseWriter, r *http.Request) {
	// parse and validate request
	var req request.Phrase
	err := parseRequest(r, &req)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	pattern, err := api.svc.GetPhrasePattern(r.Context(), req.Phrase)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respond
	respondOk(w, pattern)
}

// GetPatternExamples is handler for route POST /pattern/example
func (api *Pattern) GetPatternExamples(w http.ResponseWriter, r *http.Request) {
	// parse and validate request
	var req request.Pattern
	err := parseRequest(r, &req)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// fetch phrase examples
	examples, err := api.svc.GetPatternExamples(r.Context(), req.Tokens)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respond
	respondOk(w, examples)
}
