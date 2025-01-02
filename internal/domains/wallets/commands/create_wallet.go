package commands

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	sqlc "gitlab.rinznetwork.com/gocryptowallet/go-template/db/sqlc"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/dto"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/topics"
	kafkaClient "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

func NewCreateWalletHandler(
	log logger.Logger,
	cfg *config.Config,
	kafkaProducer kafkaClient.Producer,
	command sqlc.Querier,
	querier sqlc.Querier,
) *createWalletHandler {
	return &createWalletHandler{log: log, cfg: cfg, kafkaProducer: kafkaProducer, command: command, querier: querier}
}

func (c *createWalletHandler) Handle(ctx context.Context, createDto *dto.CreateWalletDto) (dto.CreateWalletResponseDto, error) {
	_, err := c.querier.GetWalletByUserId(ctx, createDto.UserID)
	if err != nil {
		return dto.CreateWalletResponseDto{}, errors.New("user not found")
	}

	response := dto.CreateWalletResponseDto{
		UserID:    createDto.UserID,
		BackupKey: createDto.BackupKey,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	msg, err := json.Marshal(response)

	if err != nil {
		return dto.CreateWalletResponseDto{}, err
	}

	c.kafkaProducer.PublishMessage(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topics.CREATE_WALLET_TOPIC, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
		Timestamp:      time.Now().UTC(),
	})

	return response, nil
}
