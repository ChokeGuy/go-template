package dto

import "github.com/google/uuid"

type CreateWalletDto struct {
	UserID    uuid.UUID `json:"userId" validate:"required"`
	BackupKey string    `json:"backupKey" validate:"required"`
}

type CreateWalletResponseDto struct {
	UserID    uuid.UUID `json:"userId" validate:"required"`
	BackupKey string    `json:"backupKey" validate:"required"`
	CreatedAt string    `json:"createdAt" validate:"required"`
	UpdatedAt string    `json:"updatedAt" validate:"required"`
}
