// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: wallets.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getWalletById = `-- name: GetWalletById :one
SELECT wallet_id,user_id,created_at,updated_at FROM wallets WHERE wallet_id = $1
`

type GetWalletByIdRow struct {
	WalletID  uuid.UUID          `json:"wallet_id"`
	UserID    uuid.UUID          `json:"user_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) GetWalletById(ctx context.Context, walletID uuid.UUID) (GetWalletByIdRow, error) {
	row := q.db.QueryRow(ctx, getWalletById, walletID)
	var i GetWalletByIdRow
	err := row.Scan(
		&i.WalletID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}