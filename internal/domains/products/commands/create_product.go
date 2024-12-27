package commands

import (
	"context"
	"encoding/json"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/dto"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type CreateProductCmdHandler interface {
	Handle(ctx context.Context, command *dto.CreateProductDto) error
}

type createProductHandler struct {
	log logger.Logger
	cfg *config.Config
	// kafkaProducer kafkaClient.Producer
}

func NewCreateProductHandler(log logger.Logger, cfg *config.Config) *createProductHandler {
	return &createProductHandler{log: log, cfg: cfg}
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
