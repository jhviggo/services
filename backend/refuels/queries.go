package refuels

import (
	"fmt"
	"log"

	"viggo.work/repository"
)

type Refuel struct {
	ID        int     `json:"id"`
	UserId    int     `json:"userId"`
	VehicleId int     `json:"vehicleId"`
	TotalKM   float32 `json:"totalKM"`
	TripKM    float32 `json:"tripKM"`
	Liters    float32 `json:"liters"`
	Cost      float32 `json:"cost"`
	Currency  string  `json:"currency"`
}

func fetchRefuels(userId int, vehicleId int) ([]Refuel, error) {
	var refuels []Refuel

	rows, err := repository.DB.Query(
		`SELECT id, userId, vehicleId, totalKM, tripKM, liters, cost, currency
		FROM refuels
		WHERE userId = ?
		AND vehicleId = ?`,
		userId,
		vehicleId,
	)

	if err != nil {
		return refuels, fmt.Errorf("refuels not found")
	}

	for rows.Next() {
		var myRefuel Refuel
		rows.Scan(&myRefuel.ID, &myRefuel.UserId, &myRefuel.VehicleId, &myRefuel.TotalKM, &myRefuel.TripKM, &myRefuel.Liters, &myRefuel.Cost, &myRefuel.Currency)
		refuels = append(refuels, myRefuel)
	}

	return refuels, nil
}

func insertRefuel(refuel Refuel) (Refuel, error) {
	var newRefuel Refuel

	_, err := repository.DB.Exec(
		`INSERT INTO refuels (userId, vehicleId, totalKM, tripKM, liters, cost, currency, createdAt)
		VALUES (?, ?, ?, ?, ?, ?, ?, NOW())`,
		refuel.UserId,
		refuel.VehicleId,
		refuel.TotalKM,
		refuel.TripKM,
		refuel.Liters,
		refuel.Cost,
		refuel.Currency,
	)

	if err != nil {
		return newRefuel, fmt.Errorf("unable to insert refuel")
	}

	err = repository.DB.QueryRow(
		`SELECT id, userId, vehicleId, totalKM, tripKM, liters, cost, currency
		FROM refuels
		WHERE userId = ?
		AND vehicleId = ?
		ORDER BY id DESC
		LIMIT 1`,
		refuel.UserId,
		refuel.VehicleId,
	).Scan(&newRefuel.ID, &newRefuel.UserId, &newRefuel.VehicleId, &newRefuel.TotalKM, &newRefuel.TripKM, &newRefuel.Liters, &newRefuel.Cost, &newRefuel.Currency)

	if err != nil {
		log.Println("unable to insert refuel", err.Error())
		return newRefuel, fmt.Errorf("unable to get new refuel")
	}

	return newRefuel, nil
}
