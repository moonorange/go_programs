package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Msg struct {
	ID   int32  `json:"id"`
	Body string `json:"body"`
}

func NewMsg(id int32, body string) Msg {
	return Msg{id, body}
}

// Counter represents a simple counter.
type Counter struct {
	mu    sync.Mutex
	value int32
}

// NewCounter creates and initializes a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

// Increment increments the counter value by 1.
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current value of the counter.
func (c *Counter) Value() int32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Push Message every 10 seconds to Kafka
func Publish(ctx context.Context, topic string) {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	// Create Kafka writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KAFKA_SERVER},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Publish context canceled")
			return
		case <-ticker.C:
			val := GlobalCounter.Value()
			msg := NewMsg(val, fmt.Sprintf("%s_%s_%d", topic, "test", val))
			GlobalCounter.Increment()

			// Send message to Kafka
			err := pushMsg(ctx, writer, msg, topic)
			if err != nil {
				logrus.Error("Error sending a message to Kafka: ", err)
				continue
			}

			logrus.Infof("Message sent to Kafka: %+v on Topic: %s\n", msg, topic)
		}
	}

}

func pushMsg(ctx context.Context, writer *kafka.Writer, msg Msg, topic string) error {
	val, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return writer.WriteMessages(ctx,
		kafka.Message{
			Headers: []kafka.Header{
				{Key: "Topic", Value: []byte(topic)},
			},
			Value: val,
		},
	)
}
