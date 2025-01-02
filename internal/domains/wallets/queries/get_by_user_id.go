package queries

import (
	"context"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	sqlc "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/dto"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type GetWalletByUserIdHandler interface {
	Handle(ctx context.Context, query *dto.GetWalletByUserIdDto) (*dto.WalletResponse, error)
}

type getWalletByIdHandler struct {
	log     logger.Logger
	cfg     *config.Config
	querier sqlc.Querier
}

func NewGetWalletByUserIdHandler(log logger.Logger, cfg *config.Config, querier sqlc.Querier) *getWalletByIdHandler {
	return &getWalletByIdHandler{log: log, cfg: cfg, querier: querier}
}

func (q *getWalletByIdHandler) Handle(ctx context.Context, query *dto.GetWalletByUserIdDto) (*dto.WalletResponse, error) {
	res, err := q.querier.GetWalletByUserId(ctx, query.UserID)
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
