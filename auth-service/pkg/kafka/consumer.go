package kafka

import (
	"account-service/config"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"time"
)

const AccountServiceUserCreatedTopic = "account-service.user.created"

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer() *Consumer {
	kafkaCfg := config.Config.Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     kafkaCfg.Brokers,
		GroupTopics: kafkaCfg.Consumer.Topics,
		GroupID:     kafkaCfg.Consumer.GroupID,
		StartOffset: kafka.FirstOffset,
		MaxBytes:    10e6, // 10MB
		Dialer: &kafka.Dialer{
			Timeout: 30 * time.Second,
		},
	})
	return &Consumer{
		reader: reader,
	}
}

func (kc *Consumer) StartConsuming() {
	for {
		m, err := kc.reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("error reading message", err)
			continue
		}
		log.Info().Str("topic", m.Topic).Str("partition", fmt.Sprintf("%v", m.Partition)).Str("offset", fmt.Sprintf("%v", m.Offset)).Str("key", string(m.Key)).Str("value", string(m.Value)).Msg("Message consumed")

		// process message
		switch m.Topic {
		case AccountServiceUserCreatedTopic:
			// process user created message
		}
	}
}
