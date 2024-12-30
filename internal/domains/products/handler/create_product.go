package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/dto"
	httpResponse "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/http_response"
)

// CreateProduct
// @Tags Products
// @Summary Create product
// @Description Create new product item
// @Accept json
// @Produce json
// @Success 201 {object} dto.CreateProductResponseDto
// @Router /products [post]
func (h *productsHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createDto dto.CreateProductDto
	if err := json.NewDecoder(r.Body).Decode(&createDto); err != nil {
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	if product, err := h.ps.Queries.GetProductById.Handle(h.ctx, &dto.GetProductByIdDto{ProductID: createDto.ProductID}); err != nil {
		h.log.WarnMsg("GetProductById", err)
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, err)
		return
	} else if product != nil {
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, map[string]string{"message": "Product already exists"})
		return
	}

	createDto.ProductID = uuid.New()
	if err := h.v.Struct(createDto); err != nil {
		h.log.WarnMsg("validate", err)
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, err)
		return
	}

	if err := h.ps.Commands.CreateProduct.Handle(h.ctx, &createDto); err != nil {
		h.log.WarnMsg("CreateProduct", err)
		httpResponse.ResponseWithJson(w, http.StatusBadRequest, err)
		return
	}

	httpResponse.ResponseWithJson(w, http.StatusCreated, &dto.CreateProductResponseDto{ProductID: createDto.ProductID})
}
