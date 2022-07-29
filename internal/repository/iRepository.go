package repository

import (
	"context"
	"sender/internal/entities"
)

type IRepository interface {
	WriteMessage(ctx context.Context, message entities.Message) error
}
