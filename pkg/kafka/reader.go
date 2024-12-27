package kafka

// import (
// 	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
// )

// // NewKafkaReader create new configured kafka reader
// func NewKafkaReader(kafkaURL []string, topic, groupID string) (*kafka.Consumer, error) {
// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"enable.auto.commit":           true,
// 		"auto.offset.reset":            "earliest",
// 		"queue.buffering.max.messages": queueCapacity,
// 		"session.timeout.ms":           dialTimeout,
// 		"max.poll.interval.ms":         partitionWatchInterval,
// 		"message.max.retries":          maxAttempts,
// 		"bootstrap.servers":            kafkaURL,
// 		"group.id":                     groupID,
// 		"fetch.min.bytes":              minBytes,
// 		"max.partition.fetch.bytes":    maxBytes,
// 		"heartbeat.interval.ms":        heartbeatInterval,
// 		"auto.commit.interval.ms":      commitInterval,
// 		"fetch.max.wait.ms":            maxWait,
// 		"log.connection.close":         true,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	c.SubscribeTopics([]string{topic}, nil)
// 	return c, nil
// }
