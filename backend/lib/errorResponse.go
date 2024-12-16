package lib

import (
	"encoding/json"
	"net/http"
	"services/project/models"
)

func HttpError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(models.DefaultResponse{Code: code, Message: message})
}
