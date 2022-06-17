package handler

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"simpleecommerceproductapi/product"
)

const (
	CreatePath = "/"
)

type (
	CreateService interface {
		Create(ctx context.Context, product product.Entity) (string, error)
	}

	CreateHandler struct {
		service CreateService
	}
)

func NewCreateHandler(service CreateService) CreateHandler {
	return CreateHandler{
		service: service,
	}
}

func RegisterCreateHandlers(h CreateHandler, router chi.Router) {
	router.Post(CreatePath, h.Create)
}

func (h CreateHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p product.Entity

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.service.Create(r.Context(), p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	resJSON, err := json.Marshal(map[string]string{
		"id": res,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resJSON)
	return

}
