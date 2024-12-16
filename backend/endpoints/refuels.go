package endpoints

import (
	"encoding/json"
	"net/http"
	"services/project/lib"
	"services/project/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

/* Get refuels for vehicle */
func GetRefuelsForVehicle(w http.ResponseWriter, r *http.Request) {
	vehicleId := chi.URLParam(r, "vehicleId")
	userId := chi.URLParam(r, "userId")

	vehicleOwner, err := lib.GetVehicleOwner(vehicleId)
	if vehicleOwner != userId || err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "User does not own vehicle")
		return
	}

	db := lib.GetDatabase()

	var refuels []models.Refuel
	result := db.Order("created_at DESC").Find(&refuels, "vehicle_id = ?", vehicleId)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to get refuels")
		return
	}
	err = json.NewEncoder(w).Encode(&refuels)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode data")
		return
	}
}

/* List refuels */
func ListRefuels(w http.ResponseWriter, r *http.Request) {
	var refuels []models.Refuel
	db := lib.GetDatabase()
	result := db.Find(&refuels)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to list refuels")
		return
	}
	err := json.NewEncoder(w).Encode(&refuels)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode data")
		return
	}
}

/* Add refuel */
func AddRefuel(w http.ResponseWriter, r *http.Request) {
	vehicleId := chi.URLParam(r, "vehicleId")
	userId := chi.URLParam(r, "userId")
	var refuel models.Refuel

	vehicleOwner, err := lib.GetVehicleOwner(vehicleId)
	if vehicleOwner != userId || err != nil {
		lib.HttpError(w, http.StatusBadRequest, "User does not own vehicle")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&refuel)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if refuel.TotalKM == 0 || refuel.TripKM == 0 || refuel.Liters == 0 || refuel.Cost == 0 {
		lib.HttpError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	newRefuelId := uuid.New()
	refuel.ID = newRefuelId
	vehicleIdAsUuid, err := uuid.Parse(vehicleId)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid vehicle ID")
		return
	}
	refuel.VehicleId = vehicleIdAsUuid

	db := lib.GetDatabase()
	result := db.Create(&refuel)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to save refuel")
		return
	}
	err = json.NewEncoder(w).Encode(&refuel)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode data")
		return
	}
}

/* Delete refuel */
func DeleteRefuel(w http.ResponseWriter, r *http.Request) {
	vehicleId := chi.URLParam(r, "vehicleId")
	userId := chi.URLParam(r, "userId")
	refuelId := chi.URLParam(r, "refuelId")

	vehicleOwner, err := lib.GetVehicleOwner(vehicleId)
	if vehicleOwner != userId || err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "User does not own vehicle")
		return
	}

	db := lib.GetDatabase()
	result := db.Delete(&models.Refuel{}, "vehicle_id = ? AND id = ?", vehicleId, refuelId)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to delete refuel")
		return
	}
}

/* Update refuel */
func UpdateRefuel(w http.ResponseWriter, r *http.Request) {
	vehicleId := chi.URLParam(r, "vehicleId")
	userId := chi.URLParam(r, "userId")

	vehicleOwner, err := lib.GetVehicleOwner(vehicleId)
	if vehicleOwner != userId || err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "User does not own vehicle")
		return
	}

	var refuel models.Refuel
	err = json.NewDecoder(r.Body).Decode(&refuel)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	db := lib.GetDatabase()
	result := db.Model(&models.Refuel{}).Where("id = ? AND vehicle_id = ?", refuel.ID, vehicleId).Updates(&refuel)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to update refuel")
		return
	}
	err = json.NewEncoder(w).Encode(&refuel)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to encode refuel data")
		return
	}
}
