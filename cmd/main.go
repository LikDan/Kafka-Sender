package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sender/internal/controller"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "user-messages", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	ctrl := controller.NewController(conn)
	fmt.Println(ctrl)
}
