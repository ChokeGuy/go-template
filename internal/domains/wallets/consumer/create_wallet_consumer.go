package consumer

// import (
// 	"context"
// 	"encoding/json"
// 	"time"

// 	"github.com/avast/retry-go"
// 	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
// 	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/dto"
// )

// const (
// 	retryAttempts = 3
// 	retryDelay    = 300 * time.Millisecond
// )

// var (
// 	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
// )

// func (s *WalletMessageProcessor) processCreateWallet(ctx context.Context, r *kafka.Consumer, m *kafka.Message) {
// 	var msg *dto.CreateWalletDto
// 	if err := json.Unmarshal(m.Value, &msg); err != nil {
// 		return
// 	}

// 	if err := s.v.Struct(msg); err != nil {
// 		return
// 	}

// 	if err := retry.Do(func() error {
// 		// _, err := s.querier.CreateWallet(ctx, sqlc.CreateWalletParams{
// 		// 	WalletID:   msg.WalletID,
// 		// 	Name:        msg.Name,
// 		// 	Description: msg.Description,
// 		// 	Price:       msg.Price,
// 		// })
// 		return nil
// 	}, append(retryOptions, retry.Context(ctx))...); err != nil {
// 		return
// 	}

// 	s.CommitMessage(ctx, r, m)
// }
