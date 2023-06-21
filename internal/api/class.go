package api

import (
	"github.com/go-chi/chi"
	"github.com/spongeling/admin-api/internal/auth"
	"github.com/spongeling/admin-api/internal/request"
	"github.com/spongeling/admin-api/internal/response"
	"github.com/spongeling/admin-api/internal/service"
	"log"
	"net/http"
)

type WordClass struct {
	svc *service.Service
}

func NewWordClass(svc *service.Service) *WordClass {
	return &WordClass{svc: svc}
}

func (api *WordClass) Routes(r chi.Router) {
	r.Route("/class", func(r chi.Router) {
		r.Get("/", auth.Authenticator(api.GetAllClasses))
		r.Post("/", auth.Authenticator(api.AddClass))

		r.Route("/{class_id}", func(r chi.Router) {
			r.Patch("/", auth.Authenticator(api.UpdateClass))
			r.Delete("/", auth.Authenticator(api.DeleteClass))
		})
	})
}

// GetAllClasses is handler for route GET /word/class
func (api *WordClass) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	// fetch the classes
	classes, err := api.svc.GetAllClasses(r.Context())
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respond
	respondOk(w, classes)
}

// AddClass is handler for route POST /word/class
func (api *WordClass) AddClass(w http.ResponseWriter, r *http.Request) {
	var req request.Class
	err := parseRequest(r, &req)
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	classId, err := api.svc.AddClass(r.Context(), req)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondOk(w, response.Response{
		Id:      classId,
		Status:  http.StatusOK,
		Message: "record created successfully",
	})
}

// UpdateClass is handler for route UPDATE /word/class/{class_id}
func (api *WordClass) UpdateClass(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	classId, err := extractIdFromUrl(r, "class_id")
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, "`class_id` is required")
		return
	}

	// parse and validate request
	var req request.Class
	err = parseRequest(r, &req)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// update word class
	err = api.svc.UpdateWordClass(r.Context(), classId, req)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var res = response.Response{
		Id:      classId,
		Status:  http.StatusOK,
		Message: "record updated successfully",
	}

	// respond
	respondOk(w, res)
}

// DeleteClass is handler for route DELETE /word/class/{class_id}
func (api *WordClass) DeleteClass(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	classId, err := extractIdFromUrl(r, "class_id")
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, "`class_id` is required")
		return
	}

	// delete word class
	err = api.svc.DeleteClass(r.Context(), classId)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var res = response.Response{
		Id:      classId,
		Status:  http.StatusOK,
		Message: "record deleted successfully",
	}

	// respond
	respondOk(w, res)
}
