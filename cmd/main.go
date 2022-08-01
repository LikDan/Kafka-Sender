package main

import (
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
	"sender/internal/controller"
	"sender/internal/handler"
	"sender/internal/repository"
)

func main() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "user-messages",
		Balancer: &kafka.LeastBytes{},
	}

	repo := repository.NewRepository(w)
	ctrl := controller.NewController(repo)
	handler.NewHandler(ctrl)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
