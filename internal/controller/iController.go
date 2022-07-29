package controller

import (
	"context"
	"sender/internal/entities"
)

type IController interface {
	SendMessage(ctx context.Context, message entities.Message) error
}
