package commands

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/dto"
	kafkaClient "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type CreateProductCmdHandler interface {
	Handle(ctx context.Context, command *dto.CreateProductDto) error
}

type createProductHandler struct {
	log           logger.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewCreateProductHandler(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) *createProductHandler {
	return &createProductHandler{log: log, cfg: cfg, kafkaProducer: kafkaProducer}
}

func (c *createProductHandler) Handle(ctx context.Context, command *dto.CreateProductDto) error {
	createDto := &dto.CreateProductDto{
		ProductID:   command.ProductID,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
	}

	dtoBytes, err := json.Marshal(createDto)
	if err != nil {
		return err
	}

	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		// Topic: c.cfg.KafkaTopics.ProductCreate.TopicName,
		Topic: "product-create",
		Value: dtoBytes,
		Time:  time.Now().UTC(),
	})
}
