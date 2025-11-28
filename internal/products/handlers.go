package products

import (
	"encoding/json"
	"net/http"
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

	products := []string{"musa", "tomato"}

	json.NewEncoder(w).Encode(products)
}
