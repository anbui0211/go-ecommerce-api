package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

// Curl Test
// curl -X POST "http://127.0.0.1:8999/action/stock?msg=HSG&type=MUA"

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaUrl   = "localhost:19092"
	kafkaTopic = "user_topic_vip"
)

// For producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// For consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset, // kafka.FirstOffset: Lấy tin nhắn củ
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	return &StockInfo{
		Message: msg,
		Type:    typeMsg,
	}
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
	})
}

// Consumer mua ATC
func RegisterConsumerATC(id int) {
	// group consumer
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id) // Tạo ID nhóm khác nhau cho mỗi consumer
	// kafkaGroupId:= "consumer-group-"
	reader := getKafkaReader(kafkaUrl, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer(%d) Hong Phien ATC::\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) read message error: %v\n", id, err)
		}
		fmt.Printf("Consumer(%d), hong topic %v, partition %v, offset:%v, time %d %s = %s\n",
			id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaUrl, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)
	go RegisterConsumerATC(4) // Không lấy những tin nhắn cũ

	r.Run(":8999")
}
