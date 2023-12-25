package main

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	fmt.Println("Producer started")

	topic := "helloworld"

	for i := 0; i < 10; i++ {
		fmt.Printf("send message to Kafka %d \n", (i + 1))
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte(fmt.Sprintf("Hello %d", i)),
		}

		err = producer.Produce(msg, nil)
		if err != nil {
			panic(err)
		}
	}

	producer.Flush(5 * 1000)
}
