package group

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/db"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/handler"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/service"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/middlewares"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/topics"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

func getConsumerGroupTopics() []string {
	return []string{
		topics.CREATE_WALLET_TOPIC,
	}
}

func walletGroup(router *mux.Router,
	log logger.Logger,
	mw middlewares.MiddlewareManager,
	cfg *config.Config,
	v *validator.Validate,
	ctx context.Context,
	// kafkaProducer kafkaClient.Producer,
	writerDB *db.Store,
	readerDB *db.Store,
) {
	ps := service.NewWalletService(log, cfg, writerDB, readerDB)
	// walletMessageProcessor := kafkaConsumer.NewWalletMessageProcessor(log, cfg, v)

	// log.Info("Starting Writer Kafka consumers")
	// cg := kafkaClient.NewConsumerGroup(cfg.KAFKA_BROKER[0], cfg.KAFKA_GROUP_ID, log)
	// go cg.ConsumeTopic(ctx, getConsumerGroupTopics(), kafkaConsumer.PoolSize, walletMessageProcessor.ProcessMessages)

	walletHandlers := handler.NewWalletsHandlers(router, log, mw, cfg, ps, v, ctx)
	walletHandlers.MapRoutes()
}
