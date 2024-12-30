package repository

import (
	"context"

	products "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc/products"
	users "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc/users"
	wallets "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc/wallets"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

type Repository struct {
	ProductQueries products.Querier
	WalletQueries  wallets.Querier
	UserQueries    users.Querier
}

func InitRepo(db DBTX) *Repository {
	return &Repository{
		ProductQueries: products.New(db),
		WalletQueries:  wallets.New(db),
		UserQueries:    users.New(db),
	}
}
