package commands

import (
	"context"
	"encoding/json"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/dto"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

func NewCreateWalletHandler(log logger.Logger, cfg *config.Config) *createWalletHandler {
	return &createWalletHandler{log: log, cfg: cfg}
}

func (c *createWalletHandler) Handle(ctx context.Context, command *dto.CreateWalletDto) ([]byte, error) {
	createDto := &dto.CreateWalletDto{
		WalletID:   command.WalletID,
		UserID:     command.UserID,
		PublicKey:  command.PublicKey,
		PrivateKey: command.PrivateKey,
	}

	dtoBytes, err := json.Marshal(createDto)
	if err != nil {
		return nil, err
	}

	// return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
	// 	// Topic: c.cfg.KafkaTopics.WalletCreate.TopicName,
	// 	Topic: "wallet-create",
	// 	Value: dtoBytes,
	// 	Time:  time.Now().UTC(),
	// })
	return dtoBytes, nil

}
