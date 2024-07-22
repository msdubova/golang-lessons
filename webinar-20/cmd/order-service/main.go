package main

import (
	"context"

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

	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(`{"UserID": 1, "OrderID": 3}`)}, // TODO: change to order info
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write kafka messages")
	}
}
