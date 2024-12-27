package db

import (
	"gitlab.rinznetwork.com/gocryptowallet/go-template/db/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	connPool   *pgxpool.Pool
	Repository *repository.Repository
}

func NewStore(connPool *pgxpool.Pool) *Store {
	return &Store{
		connPool:   connPool,
		Repository: repository.InitRepo(connPool),
	}
}
