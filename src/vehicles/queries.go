package vehicles

import (
	"fmt"

	"github.com/jhviggo/repository"
)

type Vehicle struct {
	ID           int    `json:"id"`
	UserId       int    `json:"userId"`
	Model        string `json:"model"`
	EstimatedKML int    `json:"estimatedKM"`
}

func insertVehicle(vehicle Vehicle) error {
	return nil
}

func fetchVehicles(userId string) ([]Vehicle, error) {
	vehicles := make([]Vehicle, 0)

	rows, err := repository.DB.Query(
		"SELECT userId, model, estimatedKML FROM vehicles WHERE userId = ?",
		userId,
	)

	if err != nil {
		return vehicles, err
	}

	for rows.Next() {
		var vehicle Vehicle
		rows.Scan(&vehicle.UserId, &vehicle.Model, &vehicle.EstimatedKML)
		vehicles = append(vehicles, vehicle)
	}

	if len(vehicles) == 0 {
		return vehicles, fmt.Errorf("user has no vehicles")
	}

	return vehicles, nil
}
