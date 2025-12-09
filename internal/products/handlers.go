package products

import (
	"log"
	"net/http"

	"github.com/Abir-Zayn/kotoNilo/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. call the service  -> List Product
	// 2. Return JSON is an HTTP Response

	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	products := []struct {
		Products []string `json:"products"`
	}{}
	json.Write(w, http.StatusOK, products)
}
