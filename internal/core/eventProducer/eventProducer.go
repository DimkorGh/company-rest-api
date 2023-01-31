package eventProducer

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type EventProducerInt interface {
	Initialize()
	Produce(message []byte, topic string)
}

type EventProducer struct {
	producer *kafka.Producer
}

func NewEventProducer() *EventProducer {
	return &EventProducer{}
}

func (ep *EventProducer) Initialize() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_ADDRESS"),
		"client.id":         os.Getenv("KAFKA_CLIENT"),
		"acks":              "all"},
	)
	if err != nil {
		log.Fatalf("Failed to create event producer: %s", err.Error())
	}

	ep.producer = producer
}

func (ep *EventProducer) Produce(message []byte, topic string) {
	deliveryChan := make(chan kafka.Event, 10000)
	err := ep.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message},
		deliveryChan,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"topic": topic,
		}).Errorf("Error while producing event: %s", err.Error())
		return
	}

	go func() {
		for e := range ep.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					logrus.Errorf("Failed to deliver message: %+v", ev.TopicPartition)
				}
			}
		}
	}()
}
