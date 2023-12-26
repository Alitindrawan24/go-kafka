package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST"),
		"group.id":          "golang",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// Wait for kafka is ready
	time.Sleep(5 * time.Second)

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
