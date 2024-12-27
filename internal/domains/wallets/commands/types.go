package commands

import (
	"context"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/dto"
	kafkaClient "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type WalletCommands struct {
	CreateWallet CreateWalletCmdHandler
}

type CreateWalletCmdHandler interface {
	Handle(ctx context.Context, command *dto.CreateWalletDto) ([]byte, error)
}

type createWalletHandler struct {
	log           logger.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}
