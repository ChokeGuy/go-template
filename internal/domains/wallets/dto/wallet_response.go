package dto

import (
	"time"
)

type WalletResponse struct {
	WalletID  string    `json:"walletId"`
	UserID    string    `json:"userId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
