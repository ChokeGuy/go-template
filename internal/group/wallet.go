package group

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/db"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/consumer"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/handler"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/service"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/middlewares"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/topics"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

func getWalletConsumerGroupTopics() []string {
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
	kafkaProducer kafka.Producer,
	writerDB *db.Store,
	readerDB *db.Store,
) {
	ws := service.NewWalletService(log, cfg, kafkaProducer, writerDB, readerDB)
	walletMessageProcessor := consumer.NewWalletMessageProcessor(log, cfg, v)

	log.Info("Starting Writer Kafka consumers")
	cg := kafka.NewConsumerGroup(cfg.KAFKA_BROKER, cfg.KAFKA_GROUP_ID, log)
	go cg.ConsumeTopic(ctx, getWalletConsumerGroupTopics(), cfg.AWS_REGION, consumer.PoolSize, walletMessageProcessor.ProcessMessages)

	walletHandlers := handler.NewWalletsHandlers(router, log, mw, cfg, ws, v, ctx)
	walletHandlers.MapRoutes()
}
