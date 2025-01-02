package repository

import (
	"context"

	wallets "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

type Repository struct {
	WalletQueries wallets.Querier
}

func InitRepo(db DBTX) *Repository {
	return &Repository{
		WalletQueries: wallets.New(db),
	}
}
