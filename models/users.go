package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Id           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName    string    `gorm:"type:varchar(30);not null"`
	LastName     string    `gorm:"type:varchar(30);not null"`
	Email        string    `gorm:"uniqueIndex;not null"`
	Password     []byte    `gorm:"not null"`
	LastedUpdate time.Time
}
