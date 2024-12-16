package models

import (
	"time"

	"github.com/google/uuid"
)

type DefaultModel struct {
	ID        uuid.UUID `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at"`
}
