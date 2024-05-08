package gorm

import (
	"time"
)

type Choosers struct {
	ID            string    `gorm:"primaryKey;not null"`
	Active        bool      `gorm:"not null"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
	DeactivatedAt time.Time `gorm:"not null"`
	Name          string    `gorm:"not null"`
	Email         string    `gorm:"unique;not null"`
	Password      string    `gorm:"not null"`
	City          string    `gorm:"not null"`
	State         string    `gorm:"not null"`
	Country       string    `gorm:"not null"`
	Day           int       `gorm:"not null"`
	Month         int       `gorm:"not null"`
	Year          int       `gorm:"not null"`
	ImageID       string    `gorm:"not null"`
}
