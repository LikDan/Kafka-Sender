package validators

import (
	"errors"
	"sender/internal/entities"
	"strings"
)

var ValidationError = errors.New("validation error")

func ValidateMessage(message entities.Message) error {
	if strings.TrimSpace(message.Message) == "" || message.Id <= 0 {
		return ValidationError
	}

	return nil
}
