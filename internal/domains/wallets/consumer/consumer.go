package consumer

// import (
// 	"context"
// 	"sync"

// 	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
// 	"github.com/go-playground/validator"
// 	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
// 	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/topics"
// 	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
// )

// const (
// 	PoolSize = 30
// )

// type WalletMessageProcessor struct {
// 	log logger.Logger
// 	cfg *config.Config
// 	v   *validator.Validate
// }

// func NewWalletMessageProcessor(log logger.Logger, cfg *config.Config, v *validator.Validate) *WalletMessageProcessor {
// 	return &WalletMessageProcessor{log: log, cfg: cfg, v: v}
// }

// func (s *WalletMessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Consumer, wg *sync.WaitGroup, workerID int) {
// 	defer wg.Done()
// 	defer r.Close()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 		}

// 		m, err := r.ReadMessage(-1)
// 		if err != nil {
// 			s.log.Warnf("workerID: %v, err: %v", workerID, err)
// 			continue
// 		}

// 		s.logProcessMessage(*m, workerID)

// 		switch *m.TopicPartition.Topic {
// 		case topics.CREATE_WALLET_TOPIC:
// 			s.processCreateWallet(ctx, r, m)
// 		}
// 	}
// }
