package service

import (
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/commands"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/queries"
	kafkaClient "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type ProductService struct {
	Commands *commands.ProductCommands
	Queries  *queries.ProductQueries
}

func NewProductService(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) *ProductService {
	createProductHandler := commands.NewCreateProductHandler(log, cfg, kafkaProducer)

	getProductByIdHandler := queries.NewGetProductByIdHandler(log, cfg)

	productCommands := commands.NewProductCommands(createProductHandler)
	productQueries := queries.NewProductQueries(getProductByIdHandler)

	return &ProductService{Commands: productCommands, Queries: productQueries}
}
