package dto

import "github.com/google/uuid"

type GetProductByIdDto struct {
	ProductID uuid.UUID `json:"productId" validate:"required,gte=0,lte=255"`
}
