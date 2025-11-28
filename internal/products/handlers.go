package products

import (
	"net/http"

	"github.com/MuhammedKasujja/ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	// return reference of the handler with this service
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListProducts(r.Context())

	if err != nil {
		// json.Write(w, http.StatusExpectationFailed, "failed to load products")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := struct {
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)
}
