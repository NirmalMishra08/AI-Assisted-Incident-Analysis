package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {

	consumer, err := sarama.NewConsumer(
		[]string{"localhost:9092"},
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(
		"system-events",
		0,
		sarama.OffsetNewest,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer partitionConsumer.Close()

	log.Println("Consumer started...")

	for msg := range partitionConsumer.Messages() {

		log.Printf(
			"Consumed: %s",
			string(msg.Value),
		)
	}
}