package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/dto"
	httpResponse "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/http_response"
)

// GetWalletByID
// @Tags Wallets
// @Summary Get product
// @Description Get product by id
// @Accept json
// @Produce json
// @Param id path string true "Wallet ID"
// @Success 200 {object} dto.WalletResponse
// @Router /wallets/{id} [get]
func (h *walletsHandlers) GetWalletByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userUUID, err := uuid.Parse(params["id"])
	if err != nil {
		h.log.WarnMsg("uuid.FromString", err)
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.ws.Queries.GetWalletByUserId.Handle(h.ctx, &dto.GetWalletByUserIdDto{UserID: userUUID})
	if err != nil {
		h.log.WarnMsg("GetWalletByUserID", err)
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	httpResponse.ResponseWithJson(w, http.StatusOK, response)
}
