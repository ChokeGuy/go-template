package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

// NewKafkaConn create new kafka connection
func NewKafkaConn(ctx context.Context, kafkaURL string) (*kafka.Conn, error) {
	return kafka.DialContext(ctx, "tcp", kafkaURL)
}
