package eventProducer

import (
	"company-rest-api/internal/core/config"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type EventProducerInt interface {
	Initialize()
	Produce(message []byte, topic string)
}

type EventProducer struct {
	producer *kafka.Producer
	cnf      *config.Config
}

func NewEventProducer(cnf *config.Config) *EventProducer {
	return &EventProducer{
		cnf: cnf,
	}
}

func (ep *EventProducer) Initialize() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": ep.cnf.Kafka.Address,
		"client.id":         ep.cnf.Kafka.Client,
		"acks":              "all",
	},
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
		Value:          message,
	},
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
