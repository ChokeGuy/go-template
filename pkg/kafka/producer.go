package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type Producer interface {
	PublishMessage(msg *kafka.Message) error
	Close()
}

type producer struct {
	log     logger.Logger
	brokers string
	w       *kafka.Producer
}

// NewProducer create new kafka producer
func NewProducer(log logger.Logger, brokers string) *producer {
	w, err := NewKafkaWriter(brokers)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}

	return &producer{log: log, brokers: brokers, w: w}
}

func (p *producer) PublishMessage(msgs *kafka.Message) error {
	deliveryChan := make(chan kafka.Event)
	err := p.w.Produce(msgs, deliveryChan)

	<-deliveryChan

	if err != nil {
		return err
	}

	return nil
}

func (p *producer) Close() {
	p.w.Close()
}
