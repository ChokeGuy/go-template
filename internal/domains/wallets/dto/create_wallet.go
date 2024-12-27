package dto

import "github.com/google/uuid"

type CreateWalletDto struct {
	WalletID   uuid.UUID `json:"walletId" validate:"required"`
	UserID     uuid.UUID `json:"userId" validate:"required"`
	PublicKey  string    `json:"publicKey" validate:"required"`
	PrivateKey string    `json:"privateKey" validate:"required"`
}

type CreateWalletResponseDto struct {
	WalletID uuid.UUID `json:"walletId,omitempty"`
	UserID   uuid.UUID `json:"userId,omitempty"`
}
