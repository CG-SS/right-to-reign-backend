package ebs

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type Replicator struct {
	Brokers []string
	Topic   string
	GroupID string
}

func getKafkaReader(brokers []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
}

func consumeFromKafka(reader *kafka.Reader) {
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

func (r *Replicator) Start() {
	reader := getKafkaReader(r.Brokers, r.Topic, r.GroupID)

	go consumeFromKafka(reader)
}
