package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/dto"
	httpResponse "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/http_response"
)

// CreateWallet
// @Tags Wallets
// @Summary Create product
// @Description Create new product item
// @Accept json
// @Produce json
// @Success 201 {object} dto.CreateWalletResponseDto
// @Router /wallets [post]
func (h *walletsHandlers) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var createDto dto.CreateWalletDto
	if err := json.NewDecoder(r.Body).Decode(&createDto); err != nil {
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	if err := h.v.Struct(createDto); err != nil {
		h.log.WarnMsg("validate", err)
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	value, err := h.ws.Commands.CreateWallet.Handle(h.ctx, &createDto)
	if err != nil {
		h.log.WarnMsg("CreateWallet", err)
		httpResponse.ResponseWithJson(w, http.StatusNotFound, err.Error())
		return
	}

	httpResponse.ResponseWithJson(w, http.StatusCreated, value)
}
