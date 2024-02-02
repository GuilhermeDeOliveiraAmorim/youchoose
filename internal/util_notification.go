package internal

import (
	"errors"
	"net/http"
)

type Notification struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewNotification(key, value string) (*Notification, error) {
	notification := &Notification{
		Key:   key,
		Value: value,
	}

	if err := notification.Validate(); err != nil {
		return nil, err
	}

	return notification, nil
}

func (n *Notification) Validate() error {
	if len(n.Key) == 0 || len(n.Key) > 50 {
		NewLogger(
			http.StatusBadRequest,
			"A chave da notificação deve ter entre 1 e 50 caracteres",
			"NewProblemDetails",
			"Entities",
			"Error",
		)

		return errors.New("a chave da notificação deve ter entre 1 e 50 caracteres")
	}

	if len(n.Value) > 255 {
		NewLogger(
			http.StatusBadRequest,
			"O valor da notificação não pode ter mais do que 255 caracteres",
			"NewProblemDetails",
			"Entities",
			"Error",
		)
		
		return errors.New("o valor da notificação não pode ter mais do que 255 caracteres")
	}

	return nil
}