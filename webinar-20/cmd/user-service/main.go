package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

const ordersTopic = "orders"

func main() {
	ctx := context.Background()

	service := &userService{
		storage: newInMemUserStorage(),
	}

	go func() {
		http.ListenAndServe(":8080", newUserHandler(service))
	}()

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   ordersTopic,
	})

	for {
		message, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to read message")
		}

		log.Info().
			Str("msg", string(message.Value)).
			Msg("Got message from kafka")

		var e orderPaidEvent

		if err := json.Unmarshal(message.Value, &e); err != nil {
			log.Warn().Err(err).Msg("Failed to decode message")
			continue
		}

		if err := service.ConsumeOrderPaidEvent(e); err != nil {
			log.Warn().Err(err).Msg("Failed to consume")
		}
	}
}
