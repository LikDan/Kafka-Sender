package controller

import (
	"context"
	"sender/internal/controller/validators"
	"sender/internal/entities"
	"sender/internal/repository"
)

type Controller struct {
	repository repository.IRepository
}

func NewController(repository repository.IRepository) *Controller {
	return &Controller{repository: repository}
}

func (c *Controller) SendMessage(ctx context.Context, message entities.Message) error {
	if err := validators.ValidateMessage(message); err != nil {
		return err
	}

	if err := c.repository.WriteMessage(ctx, message); err != nil {
		return err
	}

	return nil
}
