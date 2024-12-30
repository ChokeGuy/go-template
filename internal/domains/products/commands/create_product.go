package commands

import (
	"context"
	"encoding/json"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	sqlc "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc/products"
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
	command       sqlc.Querier
	querier       sqlc.Querier
}

func NewCreateProductHandler(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer, command sqlc.Querier, querier sqlc.Querier) *createProductHandler {
	return &createProductHandler{log: log, cfg: cfg, kafkaProducer: kafkaProducer, command: command, querier: querier}
}

func (c *createProductHandler) Handle(ctx context.Context, command *dto.CreateProductDto) error {
	createDto := &dto.CreateProductDto{
		ProductID:   command.ProductID,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
	}

	_, err := json.Marshal(createDto)
	if err != nil {
		return err
	}
	return nil
}
