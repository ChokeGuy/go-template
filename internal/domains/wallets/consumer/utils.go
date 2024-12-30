package consumer

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (s *WalletMessageProcessor) CommitMessage(ctx context.Context, r *kafka.Consumer, m *kafka.Message) {
	s.log.KafkaLogCommittedMessage(*m.TopicPartition.Topic, int(m.TopicPartition.Partition), int64(m.TopicPartition.Offset))

	if _, err := r.CommitMessage(m); err != nil {
		s.log.WarnMsg("commitMessage", err)
	}
}

func (s *WalletMessageProcessor) logProcessMessage(m kafka.Message, workerID int) {
	s.log.KafkaProcessMessage(*m.TopicPartition.Topic, int(m.TopicPartition.Partition), string(m.Value), workerID, int64(m.TopicPartition.Offset), m.Timestamp)
}

func (s *WalletMessageProcessor) commitErrMessage(ctx context.Context, r *kafka.Consumer, m *kafka.Message) {
	s.log.KafkaLogCommittedMessage(*m.TopicPartition.Topic, int(m.TopicPartition.Partition), int64(m.TopicPartition.Offset))
	if _, err := r.CommitMessage(m); err != nil {
		s.log.WarnMsg("commitMessage", err)
	}
}
