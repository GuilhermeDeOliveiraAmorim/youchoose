package internal

import (
	"time"

	"github.com/google/uuid"
)

type SharedEntity struct {
	ID            string          `json:"id"`
	Active        bool            `json:"active"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeactivatedAt time.Time       `json:"deactivated_at"`
	Notifications []Notification  `json:"notifications"`
	Errors        []ProblemDetails `json:"errors"`
}

func NewSharedEntity() *SharedEntity {
	return &SharedEntity{
		ID:            uuid.New().String(),
		Active:        true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeactivatedAt: time.Now(),
	}
}

func (se *SharedEntity) Activate() {
	se.UpdatedAt = time.Now()
	se.Active = true
}

func (se *SharedEntity) Deactivate() {
	timeNow := time.Now()
	se.DeactivatedAt = timeNow
	se.UpdatedAt = timeNow
	se.Active = false
}

func (se *SharedEntity) AddNotification(key, value string) {
	notification, _ := NewNotification(key, value)
	se.Notifications = append(se.Notifications, *notification)
}

func (se *SharedEntity) AddError(problem ProblemDetails) {
	se.Errors = append(se.Errors, problem)
}

func (se *SharedEntity) ClearNotifications() {
	se.Notifications = nil
}

func (se *SharedEntity) ClearErrors() {
	se.Errors = nil
}