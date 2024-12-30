package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/dto"
	httpResponse "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/http_response"
)

// GetProductByID
// @Tags Products
// @Summary Get product
// @Description Get product by id
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Router /products/{id} [get]
func (h *productsHandlers) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productUUID, err := uuid.Parse(params["id"])
	if err != nil {
		h.log.WarnMsg("uuid.FromString", err)
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.ps.Queries.GetProductById.Handle(h.ctx, &dto.GetProductByIdDto{ProductID: productUUID})
	if err != nil {
		h.log.WarnMsg("GetProductById", err)
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, err.Error())
		return
	}

	httpResponse.ResponseWithJson(w, http.StatusOK, response)
}
