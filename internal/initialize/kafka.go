package initialize

import (
	"ecommerce/global"
	"log"

	"github.com/segmentio/kafka-go"
)

// Init kafka producer
var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:19092"),
		Topic:    "otp-auth-topic",
		Balancer: &kafka.LeastBytes{}, // Lấy cái mới nhất
	}
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Error closing Kafka producer: %v", err)
	}
}