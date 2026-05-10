package producer

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/IBM/sarama"
)

type Event struct {
	Service   string `json:"service"`
	EventType string `json:"event_type"`
	LatencyMs int    `json:"latency_ms"`
	Timestamp string `json:"timestamp"`
}

func main() {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(
		[]string{"http://localhost:9092"},
		config,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer producer.Close()

	services := []string{
		"payment-worker",
		"email-worker",
		"api-gateway",
	}

	events := []string{
		"job_started",
		"job_completed",
		"job_failed",
	}

	for {

		event := Event{
			Service:   services[rand.Intn(len(services))],
			EventType: events[rand.Intn(len(events))],
			LatencyMs: rand.Intn(500),
			Timestamp: time.Now().Format(time.RFC3339),
		}

		jsonData, _ := json.Marshal(event)

		msg := &sarama.ProducerMessage{
			Topic: "system-events",
			Value: sarama.StringEncoder(jsonData),
		}

		_, _, err := producer.SendMessage(msg)

		if err != nil {
			log.Println(err)
		} else {
			log.Println("Produced:", string(jsonData))
		}

		time.Sleep(1 * time.Second)
	}

}
