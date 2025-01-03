package service

import (
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/db"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/commands"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/queries"
	kafkaClient "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type WalletService struct {
	Commands *commands.WalletCommands
	Queries  *queries.WalletQueries
}

func NewWalletService(
	log logger.Logger,
	cfg *config.Config,
	kafkaProducer kafkaClient.Producer,
	writerDB *db.Store,
	readerDB *db.Store,
) *WalletService {
	createWalletHandler := commands.NewCreateWalletHandler(log, cfg, kafkaProducer, writerDB.Repository.WalletQueries, readerDB.Repository.WalletQueries)
	getWalletByIdHandler := queries.NewGetWalletByUserIdHandler(log, cfg, readerDB.Repository.WalletQueries)

	walletCommands := commands.NewWalletCommands(createWalletHandler)
	walletQueries := queries.NewWalletQueries(getWalletByIdHandler)

	return &WalletService{Commands: walletCommands, Queries: walletQueries}
}
