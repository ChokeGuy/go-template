package dto

import "github.com/google/uuid"

type GetWalletByUserIdDto struct {
	UserID uuid.UUID `json:"userId" validate:"required,gte=0,lte=255"`
}
