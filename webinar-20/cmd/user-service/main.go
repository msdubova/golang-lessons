package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

const ordersTopic = "orders"

func main() {
	ctx := context.Background()

	conn, err := kafka.DialLeader(ctx, "tcp", "localhost:9092", ordersTopic, 0)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial kafka")
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	defer batch.Close()

	var message kafka.Message

	for {
		message, err = batch.ReadMessage()
		if err != nil {
			break
		}

		log.Info().
			Str("msg", string(message.Value)).
			Msg("Got message from kafka")
	}

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read batch message")
	}
}
