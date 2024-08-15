package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaReader struct {
	cons    *kafka.Consumer
	topic   string
	sigchan chan os.Signal
}

func NewKafkaReader(topic string) *KafkaReader {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		// User-specific properties that you must set
		"bootstrap.servers": "localhost:9092",
		// Fixed properties
		"group.id":          "v2_streams",
		"auto.offset.reset": "earliest"})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	return &KafkaReader{
		cons:    c,
		topic:   topic,
		sigchan: sigchan,
	}
}

func (c *KafkaReader) Read(out chan []byte) {
	run := true
	for run {
		select {
		case sig := <-c.sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.cons.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			out <- ev.Value
		}
	}
	c.cons.Close()

}
