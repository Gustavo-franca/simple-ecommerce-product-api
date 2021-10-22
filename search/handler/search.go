package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

const (
	getByIDPath      = "/{id}"
	getByFiltersPath = "/"
)

type SearchHandler struct {
}

func New() SearchHandler {
	return SearchHandler{}
}

func RegisterWebHandlers(h SearchHandler, router chi.Router) {
	router.Get(getByIDPath, h.GetByID)
	router.Get(getByFiltersPath, h.GetByFilters)
}
func (h SearchHandler) GetByFilters(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusAccepted)
}
func (h SearchHandler) GetByID(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusAccepted)
}
