package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewKafkaWriter(brokers string, region string, accessKey, secretKey string) (*kafka.Producer, error) {
	tokenProvider, err := NewIAMTokenProvider(
		region,
		accessKey,
		secretKey,
	)
	if err != nil {
		return nil, err
	}

	token, err := tokenProvider.Token()

	if err != nil {
		return nil, err
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"security.protocol":  "SASL_SSL",
		"sasl.mechanisms":    "OAUTHBEARER",
		"acks":               writerRequiredAcks,
		"retries":            writerMaxAttempts,
		"compression.type":   "snappy",
		"socket.timeout.ms":  int(writerReadTimeout.Milliseconds()),
		"request.timeout.ms": int(writerWriteTimeout.Milliseconds()),
		"linger.ms":          0,
	})

	if err != nil {
		return nil, err
	}

	p.SetOAuthBearerToken(token)
	return p, nil
}
