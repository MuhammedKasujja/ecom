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
	// 1. Call the service -> ListProduct
	// 2. Return JSON in the HTTP response

	products := struct {
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)
}
