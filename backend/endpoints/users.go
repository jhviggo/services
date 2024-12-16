package endpoints

import (
	"encoding/json"
	"net/http"
	"services/project/lib"
	"services/project/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

/* Get users */
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	userId := chi.URLParam(r, "userId")
	db := lib.GetDatabase()

	uuid, err := uuid.Parse(userId)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	result := db.First(&user, "id = ?", uuid)
	if result.Error != nil {
		lib.HttpError(w, http.StatusNotFound, "User not found")
		return
	}
	user.Password = ""
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode data")
		return
	}
}

/* List users */
func ListUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db := lib.GetDatabase()
	result := db.Find(&users)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to list users")
		return
	}
	for i := 0; i < len(users); i++ {
		users[i].Password = ""
	}

	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode data")
		return
	}
}

/* Add user */
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := uuid.New()
	user.ID = id

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	hash, err := lib.HashPasswordWithSalt(user.Password)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to hash password")
		return
	}
	user.Password = hash

	db := lib.GetDatabase()
	result := db.Create(&user)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to add new user")
		return
	}

	user.Password = ""
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode data")
		return
	}
}

/* Delete user */
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid user id")
		return
	}
	db := lib.GetDatabase()
	result := db.Delete(&models.User{}, "id = ?", userUuid)
	if result.Error != nil || result.RowsAffected == 0 {
		lib.HttpError(w, http.StatusBadRequest, "Unable to remove user")
		return
	}
}

/* Update user */
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user.Password = ""
	uuid, err := uuid.Parse(userId)
	if err != nil {
		lib.HttpError(w, http.StatusBadRequest, "Invalid user id")
		return
	}
	user.ID = uuid
	user.Password = ""

	db := lib.GetDatabase()

	result := db.Model(&models.User{}).Where("id = ?", userId).Updates(&user)
	if result.Error != nil {
		lib.HttpError(w, http.StatusBadRequest, "Unable to update user")
		return
	}

	result = db.First(&user, "id = ?", userId)
	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to get user after update")
		return
	}
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Failed to encode data")
		return
	}
}
