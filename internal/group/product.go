package group

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/db"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/consumer"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/handler"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/service"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/middlewares"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/topics"
	kafkaClient "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

func getProductConsumerGroupTopics() []string {
	return []string{
		topics.CREATE_PRODUCT_TOPIC,
	}
}

func productGroup(router *mux.Router,
	log logger.Logger,
	mw middlewares.MiddlewareManager,
	cfg *config.Config,
	v *validator.Validate,
	ctx context.Context,
	kafkaProducer kafkaClient.Producer,
	writerDB *db.Store,
	readerDB *db.Store,
) {
	ps := service.NewProductService(log, cfg, kafkaProducer, writerDB, readerDB)
	productMessageProcessor := consumer.NewProductMessageProcessor(log, cfg, v)

	log.Info("Starting Writer Kafka consumers")
	cg := kafkaClient.NewConsumerGroup(cfg.KAFKA_BROKER, cfg.KAFKA_GROUP_ID, log)
	go cg.ConsumeTopic(ctx, getProductConsumerGroupTopics(), consumer.PoolSize, productMessageProcessor.ProcessMessages)

	productHandlers := handler.NewProductsHandlers(router, log, mw, cfg, ps, v, ctx)
	productHandlers.MapRoutes()
}
