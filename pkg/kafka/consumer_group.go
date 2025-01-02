package kafka

import (
	"context"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

// MessageProcessor processor methods must implement kafka.Worker func method interface
type MessageProcessor interface {
	ProcessMessages(ctx context.Context, r *kafka.Consumer, wg *sync.WaitGroup, workerID int)
}

// Worker kafka consumer worker fetch and process messages from reader
type Worker func(ctx context.Context, r *kafka.Consumer, wg *sync.WaitGroup, workerID int)

type ConsumerGroup interface {
	ConsumeTopic(ctx context.Context, cancel context.CancelFunc, groupID, topic string, poolSize int, worker Worker)
	GetNewKafkaReader(kafkaURL []string, topic, groupID string) *kafka.Consumer
	GetNewKafkaWriter(topic string) *kafka.Producer
}

type consumerGroup struct {
	Brokers string
	GroupID string
	log     logger.Logger
}

// NewConsumerGroup kafka consumer group constructor
func NewConsumerGroup(brokers string, groupID string, log logger.Logger) *consumerGroup {
	return &consumerGroup{Brokers: brokers, GroupID: groupID, log: log}
}

// GetNewKafkaReader create new kafka reader
func (c *consumerGroup) GetNewKafkaReader(kafkaURL string, groupTopics []string, groupID string, region string) (*kafka.Consumer, error) {
	return NewKafkaReader(kafkaURL, groupTopics, groupID, region)
}

// GetNewKafkaWriter create new kafka producer
func (c *consumerGroup) GetNewKafkaWriter(region, accessKey, secretKey string) (*kafka.Producer, error) {
	return NewKafkaWriter(c.Brokers, region, accessKey, secretKey)
}

// ConsumeTopic start consumer group with given worker and pool size
func (c *consumerGroup) ConsumeTopic(ctx context.Context, groupTopics []string, region string, poolSize int, worker Worker) {
	r, err := c.GetNewKafkaReader(c.Brokers, groupTopics, c.GroupID, region)

	if err != nil {
		return
	}

	defer func() {
		if err := r.Close(); err != nil {
			c.log.Warnf("consumerGroup.r.Close: %v", err)
		}
	}()

	c.log.Infof("Starting consumer groupID: %s, topic: %+v, pool size: %v", c.GroupID, groupTopics, poolSize)

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, wg, i)
	}
	wg.Wait()
}
