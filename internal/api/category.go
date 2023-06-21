package api

import (
	"github.com/go-chi/chi"
	"github.com/spongeling/admin-api/internal/auth"
	"github.com/spongeling/admin-api/internal/errors"
	"github.com/spongeling/admin-api/internal/request"
	"github.com/spongeling/admin-api/internal/response"
	"github.com/spongeling/admin-api/internal/service"
	"log"
	"net/http"
)

type Category struct {
	svc *service.Service
}

func NewCategory(svc *service.Service) *Category {
	return &Category{svc: svc}
}

func (api *Category) Routes(r chi.Router) {
	r.Route("/category", func(r chi.Router) {
		r.Post("/", auth.Authenticator(api.AddCategory))

		r.Get("/top", auth.Authenticator(api.GetAllTopLevelCategories))

		r.Route("/{category_id}", func(r chi.Router) {
			r.Get("/", auth.Authenticator(api.GetCategory))
			r.Patch("/", auth.Authenticator(api.UpdateCategory))
			r.Delete("/", auth.Authenticator(api.DeleteCategory))

			r.Get("/info", auth.Authenticator(api.GetCategoryInfo))
			r.Get("/subcategories", auth.Authenticator(api.GetSubCategories))
		})
	})
}

// GetAllTopLevelCategories is a handler for GET /category/top
func (api *Category) GetAllTopLevelCategories(w http.ResponseWriter, r *http.Request) {
	// fetch categories from category service
	categories, err := api.svc.GetAllTopLevelCategories(r.Context())
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respond
	respondOk(w, categories)
}

// GetCategoryInfo is a handler for GET /category/{category_id}
func (api *Category) GetCategoryInfo(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	categoryId, err := extractIdFromUrl(r, "category_id")
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, "`category_id` is required")
		return
	}

	// fetch category details
	categoryInfo, err := api.svc.GetCategoryInfo(r.Context(), categoryId)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respond
	respondOk(w, categoryInfo)
}

// GetSubCategories is a handler for GET /category/{category_id}/subcategories
func (api *Category) GetSubCategories(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	categoryId, err := extractIdFromUrl(r, "category_id")
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, "`category_id` is required")
		return
	}

	// fetch categories from category service
	categories, err := api.svc.GetSubCategories(r.Context(), categoryId)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respond
	respondOk(w, categories)
}

// AddCategory is handler for route POST /category
func (api *Category) AddCategory(w http.ResponseWriter, r *http.Request) {
	// parse and validate request
	var req request.CategoryDetails
	err := parseRequest(r, &req)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// save phrase pattern
	id, err := api.svc.AddCategory(r.Context(), req)
	if err != nil {
		log.Println(err)
		if e, ok := err.(*errors.Error); ok {
			respondErrorMessage(w, e.ErrorType(), e.Error())
		} else {
			respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// build response
	var res = response.Response{
		Id:      id,
		Status:  http.StatusOK,
		Message: "record created successfully",
	}

	// respond
	respondOk(w, res)
}

// UpdateCategory is handler for route PATCH /category/{category_id}
func (api *Category) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	categoryId, err := extractIdFromUrl(r, "category_id")
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, "`phrase_id` is required")
		return
	}

	// parse and validate request
	var req request.UpdatePatternPhrase
	err = parseRequest(r, &req)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// update phrase pattern
	err = api.svc.UpdateCategory(r.Context(), categoryId, req)
	if err != nil {
		log.Println(err)
		if e, ok := err.(*errors.Error); ok {
			respondErrorMessage(w, e.ErrorType(), e.Error())
		} else {
			respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// build response
	var res = response.Response{
		Id:      categoryId,
		Status:  http.StatusOK,
		Message: "record updated successfully",
	}

	// respond
	respondOk(w, res)
}

// DeleteCategory is handler for route DELETE /category/{category_id}
func (api *Category) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	categoryId, err := extractIdFromUrl(r, "category_id")
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, "`phrase_id` is required")
		return
	}

	// delete phrase pattern
	err = api.svc.DeleteCategory(r.Context(), categoryId)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var res = response.Response{
		Id:      categoryId,
		Status:  http.StatusOK,
		Message: "record deleted successfully",
	}

	// respond
	respondOk(w, res)
}

// GetCategory is a handler for GET /category/{category_id}
func (api *Category) GetCategory(w http.ResponseWriter, r *http.Request) {
	// parse url parameter
	categoryId, err := extractIdFromUrl(r, "category_id")
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, "`category_id` is required")
		return
	}

	// fetch category pos example from category service
	example, err := api.svc.GetCategory(r.Context(), categoryId)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respond
	respondOk(w, example)
}
