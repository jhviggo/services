package models

import (
	"time"

	"github.com/google/uuid"
)

type Refuel struct {
	DefaultModel
	VehicleId   uuid.UUID `json:"vehicle_id,omitempty" gorm:"index"`
	TotalKM     float32   `json:"total_km"`
	TripKM      float32   `json:"trip_km"`
	Liters      float32   `json:"liters,omitempty"`
	Cost        float32   `json:"cost,omitempty"`
	Currency    string    `json:"currency,omitempty"`
	Coordinates string    `json:"coordinates"`
	FuelDate    time.Time `json:"fuel_date"`
}
