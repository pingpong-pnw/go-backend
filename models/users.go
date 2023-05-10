package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Id           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName    string    `json:"firstname" binding:"required,max=30" gorm:"type:varchar(30);not null"`
	LastName     string    `json:"lastname" binding:"required,max=30" gorm:"type:varchar(30);not null"`
	Email        string    `json:"email" binding:"required,email" gorm:"uniqueIndex;not null"`
	Password     string    `json:"password" binding:"required" gorm:"not null"`
	LastedUpdate time.Time
}
