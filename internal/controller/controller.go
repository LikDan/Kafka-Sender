package controller

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"sender/internal/controller/validators"
	"sender/internal/entities"
)

type Controller struct {
	connection *kafka.Conn
}

func NewController(connection *kafka.Conn) *Controller {
	return &Controller{connection: connection}
}

func (c *Controller) SendMessage(_ context.Context, message entities.Message) error {
	if err := validators.ValidateMessage(message); err != nil {
		return err
	}

	bytesMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	kafkaMessage := kafka.Message{Value: bytesMessage}
	if _, err = c.connection.WriteMessages(kafkaMessage); err != nil {
		return err
	}

	return nil
}
