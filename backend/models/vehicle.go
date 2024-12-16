package models

import "github.com/google/uuid"

type Vehicle struct {
	DefaultModel
	UserId       uuid.UUID `json:"user_id" gorm:"index"`
	Model        string    `json:"model"`
	EstimatedKML int       `json:"estimated_kml"`
	Refuel       []Refuel  `json:",omitempty" gorm:"foreignKey:VehicleId;references:ID"`
}
