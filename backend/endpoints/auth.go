package endpoints

import (
	"encoding/json"
	"net/http"
	"services/project/lib"
	"services/project/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	json.NewDecoder(r.Body).Decode(&credentials)

	username := credentials.Username
	password := credentials.Password

	db := lib.GetDatabase()
	var user models.User
	result := db.First(&user, "username = ?", username)
	if result.Error != nil {
		lib.HttpError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err := lib.ValidatePassword(user.Password, password)
	if err != nil {
		lib.HttpError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token, err := lib.CreateToken(user.ID.String(), user.Role)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	user.Password = ""
	user.Token = token
	json.NewEncoder(w).Encode(&user)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	// @TODO implement
}
