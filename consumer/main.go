package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "golang",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	fmt.Println("Consumer started")

	err = consumer.Subscribe("helloworld", nil)
	if err != nil {
		panic(err)
	}

	for {
		message, err := consumer.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Received message: %s\n", message.Value)
		}
	}
}
