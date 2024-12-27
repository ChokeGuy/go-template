package handler

import (
	"net/http"
)

func (h *walletsHandlers) MapRoutes() {
	s := h.router.PathPrefix("/wallet").Subrouter()
	s.HandleFunc("", h.CreateWallet).Methods(http.MethodPost)
	s.HandleFunc("/{id}", h.GetWalletByID).Methods(http.MethodGet)
}
