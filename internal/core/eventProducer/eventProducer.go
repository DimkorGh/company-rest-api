package eventProducer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"

	"company-rest-api/internal/core/config"
	"company-rest-api/internal/core/log"
)

type EventProducerInt interface {
	Initialize()
	Produce(message []byte, topic string)
}

type EventProducer struct {
	cnf      *config.Config
	logger   log.LoggerInt
	producer *kafka.Producer
}

func NewEventProducer(cnf *config.Config, logger log.LoggerInt) *EventProducer {
	return &EventProducer{
		cnf:    cnf,
		logger: logger,
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
		ep.logger.Fatalf("Failed to create event producer: %s", err.Error())
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
		ep.logger.Errorw(
			"Error while producing event",
			zap.String("topic", topic),
			zap.String("error", err.Error()),
		)

		return
	}

	go func() {
		for e := range ep.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					ep.logger.Errorf("Failed to deliver message: %+v", ev.TopicPartition)
				}
			}
		}
	}()
}
