package vehicles

import (
	"fmt"

	"github.com/jhviggo/repository"
)

type Vehicle struct {
	ID           int    `json:"id"`
	UserId       int    `json:"userId"`
	Model        string `json:"model"`
	EstimatedKML int    `json:"estimatedKML"`
}

func insertVehicle(vehicle Vehicle) (Vehicle, error) {
	var newVehicle Vehicle

	_, err := repository.DB.Exec(
		`INSERT INTO vehicles (userId, model, estimatedKML, createdAt)
		VALUES (?, ?, ?, NOW());`,
		vehicle.UserId,
		vehicle.Model,
		vehicle.EstimatedKML,
	)

	if err != nil {
		return vehicle, fmt.Errorf("unable to insert vehicle")
	}

	err = repository.DB.QueryRow(
		`SELECT id, userId, model, estimatedKML FROM vehicles
		ORDER BY id DESC
		LIMIT 1;`,
	).Scan(&newVehicle.ID, &newVehicle.UserId, &newVehicle.Model, &newVehicle.EstimatedKML)

	if err != nil {
		return newVehicle, fmt.Errorf("unable to get newly added vehicle")
	}

	return newVehicle, nil
}

func fetchVehicles(userId string) ([]Vehicle, error) {
	vehicles := make([]Vehicle, 0)

	rows, err := repository.DB.Query(
		"SELECT id, userId, model, estimatedKML FROM vehicles WHERE userId = ?",
		userId,
	)

	if err != nil {
		return vehicles, err
	}

	for rows.Next() {
		var vehicle Vehicle
		rows.Scan(&vehicle.ID, &vehicle.UserId, &vehicle.Model, &vehicle.EstimatedKML)
		vehicles = append(vehicles, vehicle)
	}

	if len(vehicles) == 0 {
		return vehicles, fmt.Errorf("user has no vehicles")
	}

	return vehicles, nil
}
