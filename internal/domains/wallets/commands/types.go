package commands

import (
	"context"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	sqlc "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/dto"
	kafkaClient "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type WalletCommands struct {
	CreateWallet CreateWalletCmdHandler
}

type CreateWalletCmdHandler interface {
	Handle(ctx context.Context, command *dto.CreateWalletDto) (dto.CreateWalletResponseDto, error)
}

type createWalletHandler struct {
	log           logger.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
	command       sqlc.Querier
	querier       sqlc.Querier
}
