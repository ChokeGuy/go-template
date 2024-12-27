package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/avast/retry-go"
	"github.com/segmentio/kafka-go"
	sqlc "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc/products"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/dto"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

func (s *ProductMessageProcessor) processCreateProduct(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	var msg *dto.CreateProductDto
	if err := json.Unmarshal(m.Value, &msg); err != nil {
		return
	}

	if err := s.v.Struct(msg); err != nil {
		return
	}

	if err := retry.Do(func() error {
		_, err := s.querier.CreateProduct(ctx, sqlc.CreateProductParams{
			ProductID:   msg.ProductID,
			Name:        msg.Name,
			Description: msg.Description,
			Price:       msg.Price,
		})
		return err
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		return
	}

	s.CommitMessage(ctx, r, m)
}