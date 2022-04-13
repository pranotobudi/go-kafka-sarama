package producer

import (
	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func (p *KafkaProducer) SendMessage(topic, msg string) error {
	return nil
}
