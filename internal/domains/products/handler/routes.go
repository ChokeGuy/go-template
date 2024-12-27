package handler

import (
	"net/http"
)

func (h *productsHandlers) MapRoutes() {
	s := h.router.PathPrefix("/product").Subrouter()
	s.HandleFunc("", h.CreateProduct).Methods(http.MethodPost)
	s.HandleFunc("/{id}", h.GetProductByID).Methods(http.MethodGet)
}
