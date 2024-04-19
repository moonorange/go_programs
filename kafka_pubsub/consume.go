package main

import (
	"context"
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	Handle(msg kafka.Message)
}

func Consume(ctx context.Context, topic string, handler Handler) {
	defer func() {
		logrus.Infof("stop consuming %s from kafka %s", topic, KAFKA_SERVER)
		wg.Done()
	}()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{KAFKA_SERVER},
		GroupID: "test_consumer",
		Topic:   topic,
	})

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil && strings.Contains(err.Error(), "context canceled") {
			return
		}
		if err != nil {
			logrus.Errorln("Error reading message from Kafka: ", err)
			continue
		}
		handler.Handle(msg)

		// Process message
		logrus.Infof("Received message: %s on Topic: %s\n", string(msg.Value), topic)
	}
}
