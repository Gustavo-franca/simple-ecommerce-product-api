package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
	"simpleecommerceproductapi/product"
	"simpleecommerceproductapi/product/search"
)

const (
	getByIDPath      = "/{id}"
	getByFiltersPath = "/"
)

type (
	Service interface {
		SearchByFilters(params search.Params) ([]product.Entity, error)
		SearchByID(id string) (product.Entity, error)
	}

	SearchHandler struct {
		service Service
	}
)

func New(service Service) SearchHandler {
	return SearchHandler{
		service: service,
	}
}

func RegisterWebHandlers(h SearchHandler, router chi.Router) {
	router.Get(getByIDPath, h.GetByID)
	router.Get(getByFiltersPath, h.GetByFilters)
}
func (h SearchHandler) GetByFilters(w http.ResponseWriter, r *http.Request) {
	params, err := transformToParams(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.service.SearchByFilters(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resJSON)
	return

}

func (h SearchHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.service.SearchByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resJSON)
	return
}

func transformToParams(query url.Values) (search.Params, error) {
	return search.Params{
		Description: query.Get("description"),
		Title:       query.Get("title"),
	}, nil
}
