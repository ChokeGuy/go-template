package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// NewWriter create new configured kafka writer
func NewKafkaWriter(brokers string) (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"acks":               writerRequiredAcks,
		"retries":            writerMaxAttempts,
		"compression.type":   "snappy",
		"socket.timeout.ms":  int(writerReadTimeout.Milliseconds()),
		"request.timeout.ms": int(writerWriteTimeout.Milliseconds()),
		"linger.ms":          0,
	})
}
