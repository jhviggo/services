package endpoints

import (
	"encoding/json"
	"net/http"
	"services/project/lib"
	"services/project/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

/* Get vehicles for user */
func GetVehiclesForUser(w http.ResponseWriter, r *http.Request) {
	var vehicles []models.Vehicle
	userId := chi.URLParam(r, "userId")
	db := lib.GetDatabase()

	uuid, err := uuid.Parse(userId)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	result := db.Find(&vehicles, "user_id = ?", uuid)

	if result.Error != nil {
		lib.HttpError(w, http.StatusNotFound, "Vehicle not found")
		return
	}
	err = json.NewEncoder(w).Encode(&vehicles)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode data")
		return
	}
}

/* List vehicles */
func ListVehicles(w http.ResponseWriter, r *http.Request) {
	var vehicles []models.Vehicle
	db := lib.GetDatabase()
	result := db.Find(&vehicles)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to list vehicles")
		return
	}
	err := json.NewEncoder(w).Encode(&vehicles)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode data")
		return
	}
}

/* Add vehicle */
func AddVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle models.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	userId := chi.URLParam(r, "userId")
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid vehicle id")
		return
	}

	newId := uuid.New()
	vehicle.ID = newId
	vehicle.UserId = userUuid

	db := lib.GetDatabase()
	result := db.Create(&vehicle)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to add new vehicle")
		return
	}
	err = json.NewEncoder(w).Encode(&vehicle)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode data")
		return
	}
}

/* Delete vehicle */
func DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid vehicle id")
		return
	}

	vehicleId := chi.URLParam(r, "vehicleId")
	vehicleUuid, err := uuid.Parse(vehicleId)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid vehicle id")
		return
	}

	db := lib.GetDatabase()
	result := db.Delete(&models.Vehicle{}, "id = ? AND user_id = ?", vehicleUuid, userUuid)
	if result.Error != nil || result.RowsAffected == 0 {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to remove vehicle")
		return
	}
}

/* Update vehicle */
func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	_, err := uuid.Parse(userId)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid vehicle id")
		return
	}

	vehicleId := chi.URLParam(r, "vehicleId")
	_, err = uuid.Parse(vehicleId)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid vehicle id")
		return
	}

	var vehicle models.Vehicle
	err = json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	db := lib.GetDatabase()

	result := db.Model(&models.Vehicle{}).Where("id = ? AND user_id = ?", vehicleId, userId).Updates(&vehicle)
	if result.Error != nil {
		lib.HttpError(w, http.StatusBadRequest, "Unable to update vehicle")
		return
	}

	result = db.First(&vehicle, "id = ? AND user_id = ?", vehicleId, userId)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to get vehicle after update")
		return
	}
	err = json.NewEncoder(w).Encode(&vehicle)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode vehicle data")
		return
	}
}
