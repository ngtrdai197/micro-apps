package kafka

import (
	"account-service/config"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const AccountServiceUserCreatedTopic = "account-service.user.created"

type Consumer struct {
	reader *kafka.Reader
	stopCh chan struct{}
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
		stopCh: make(chan struct{}),
	}
}

func (kc *Consumer) StartConsuming() {
	log.Info().Msg("Started consuming messages")

	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
		<-sigterm
		log.Info().Msg("Termination signal received, stopping consumer...")
		close(kc.stopCh)
	}()

	for {
		select {
		case <-kc.stopCh:
			log.Info().Msg("Stopping consumer...")
			err := kc.reader.Close()
			if err != nil {
				log.Error().Err(err).Msg("Error closing reader")
				return
			}
			log.Info().Msg("Consumer stopped successfully")
			return
		default:
			m, err := kc.reader.ReadMessage(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("Error reading message")
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
}
