package service

import (
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/commands"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/queries"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type WalletService struct {
	Commands *commands.WalletCommands
	Queries  *queries.WalletQueries
}

func NewWalletService(log logger.Logger, cfg *config.Config) *WalletService {
	createWalletHandler := commands.NewCreateWalletHandler(log, cfg)

	getWalletByIdHandler := queries.NewGetWalletByIdHandler(log, cfg)

	walletCommands := commands.NewWalletCommands(createWalletHandler)
	walletQueries := queries.NewWalletQueries(getWalletByIdHandler)

	return &WalletService{Commands: walletCommands, Queries: walletQueries}
}
