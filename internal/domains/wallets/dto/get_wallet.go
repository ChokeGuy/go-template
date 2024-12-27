package dto

import "github.com/google/uuid"

type GetWalletByIdDto struct {
	WalletID uuid.UUID `json:"walletId" validate:"required,gte=0,lte=255"`
}
