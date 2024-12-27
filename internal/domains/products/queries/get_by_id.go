package queries

import (
	"context"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/db/repository"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/dto"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type GetProductByIdHandler interface {
	Handle(ctx context.Context, query *dto.GetProductByIdDto) (*dto.ProductResponse, error)
}

type getProductByIdHandler struct {
	log        logger.Logger
	cfg        *config.Config
	repository *repository.Repository
}

func NewGetProductByIdHandler(log logger.Logger, cfg *config.Config) *getProductByIdHandler {
	return &getProductByIdHandler{log: log, cfg: cfg}
}

func (q *getProductByIdHandler) Handle(ctx context.Context, query *dto.GetProductByIdDto) (*dto.ProductResponse, error) {
	res, err := q.repository.ProductQueries.GetProductById(ctx, query.ProductID)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ProductID:   res.ProductID.String(),
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		CreatedAt:   res.CreatedAt.Time,
		UpdatedAt:   res.UpdatedAt.Time,
	}, nil
}
