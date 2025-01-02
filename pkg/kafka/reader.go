package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// NewKafkaReader create new configured kafka reader
func NewKafkaReader(kafkaURL string, topics []string, groupID string, region string) (*kafka.Consumer, error) {
	tokenProvider := &IAMTokenProvider{region: region}
	token, err := tokenProvider.Token()

	if err != nil {
		return nil, err
	}
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"security.protocol":            "SASL_SSL",
		"sasl.mechanism":               "OAUTHBEARER",
		"enable.auto.commit":           true,
		"auto.offset.reset":            "earliest",
		"queue.buffering.max.messages": queueCapacity,
		"session.timeout.ms":           dialTimeout,
		"max.poll.interval.ms":         partitionWatchInterval,
		"message.max.retries":          maxAttempts,
		"bootstrap.servers":            kafkaURL,
		"group.id":                     groupID,
		"fetch.min.bytes":              minBytes,
		"max.partition.fetch.bytes":    maxBytes,
		"heartbeat.interval.ms":        heartbeatInterval,
		"auto.commit.interval.ms":      commitInterval,
		"fetch.max.wait.ms":            maxWait,
		"log.connection.close":         true,
	})

	if err != nil {
		return nil, err
	}

	c.SetOAuthBearerToken(token)
	c.SubscribeTopics(topics, nil)
	return c, nil
}
