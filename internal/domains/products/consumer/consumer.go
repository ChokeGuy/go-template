package consumer

import (
	"context"
	"sync"

	"github.com/go-playground/validator"
	"github.com/segmentio/kafka-go"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	sqlc "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc/products"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

const (
	PoolSize = 30
)

type ProductMessageProcessor struct {
	log     logger.Logger
	cfg     *config.Config
	v       *validator.Validate
	querier sqlc.Querier
}

func NewProductMessageProcessor(log logger.Logger, cfg *config.Config, v *validator.Validate) *ProductMessageProcessor {
	return &ProductMessageProcessor{log: log, cfg: cfg, v: v}
}

func (s *ProductMessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err := r.FetchMessage(ctx)
		if err != nil {
			s.log.Warnf("workerID: %v, err: %v", workerID, err)
			continue
		}

		s.logProcessMessage(m, workerID)

		switch m.Topic {
		// case s.cfg.KafkaTopics.ProductCreate.TopicName:
		// s.processCreateProduct(ctx, r, m)
		case "product-create":
			s.processCreateProduct(ctx, r, m)
		}
	}
}
