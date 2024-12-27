package commands

import (
	"context"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	sqlc "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc/wallets"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/dto"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type WalletCommands struct {
	CreateWallet CreateWalletCmdHandler
}

type CreateWalletCmdHandler interface {
	Handle(ctx context.Context, command *dto.CreateWalletDto) (dto.CreateWalletResponseDto, error)
}

type createWalletHandler struct {
	log     logger.Logger
	cfg     *config.Config
	command sqlc.Querier
	querier sqlc.Querier
	// kafkaProducer kafkaClient.Producer
}
