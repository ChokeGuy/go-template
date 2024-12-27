package queries

import (
	"context"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/db/repository"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/dto"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type GetWalletByIdHandler interface {
	Handle(ctx context.Context, query *dto.GetWalletByIdDto) (*dto.WalletResponse, error)
}

type getWalletByIdHandler struct {
	log        logger.Logger
	cfg        *config.Config
	repository *repository.Repository
}

func NewGetWalletByIdHandler(log logger.Logger, cfg *config.Config) *getWalletByIdHandler {
	return &getWalletByIdHandler{log: log, cfg: cfg}
}

func (q *getWalletByIdHandler) Handle(ctx context.Context, query *dto.GetWalletByIdDto) (*dto.WalletResponse, error) {
	res, err := q.repository.WalletQueries.GetWalletById(ctx, query.WalletID)
	if err != nil {
		return nil, err
	}

	return &dto.WalletResponse{
		WalletID:  res.WalletID.String(),
		UserID:    res.UserID.String(),
		CreatedAt: res.CreatedAt.Time,
		UpdatedAt: res.UpdatedAt.Time,
	}, nil
}
