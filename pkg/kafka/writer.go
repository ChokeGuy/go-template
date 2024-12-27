package kafka

// import (
// 	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
// )

// // NewWriter create new configured kafka writer
// func NewWriter(brokers string) (*kafka.Producer, error) {
// 	return kafka.NewProducer(&kafka.ConfigMap{
// 		"bootstrap.servers":  brokers,
// 		"acks":               writerRequiredAcks,
// 		"retries":            writerMaxAttempts,
// 		"compression.type":   "snappy",
// 		"socket.timeout.ms":  writerReadTimeout.Milliseconds(),
// 		"request.timeout.ms": writerWriteTimeout.Milliseconds(),
// 		"linger.ms":          0,
// 	})
// }
