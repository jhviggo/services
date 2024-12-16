package endpoints

import (
	"encoding/json"
	"net/http"
	"services/project/lib"
	"services/project/models"
	"time"

	"github.com/patrickmn/go-cache"
)

var VIGGO_USERNAME string = "viggosdetour"

// response cache for viggosdetour user refuels
var refuelsCache *cache.Cache = cache.New(60*time.Minute, 90*time.Minute)

func ListViggosdetourRefuels(w http.ResponseWriter, _ *http.Request) {
	var refuels []models.Refuel
	db := lib.GetDatabase()

	cachedRefuels, found := refuelsCache.Get(VIGGO_USERNAME)
	if !found {
		result := db.Raw(`
			SELECT id, created_at, updated_at, total_km, trip_km, coordinates
			FROM refuels
			WHERE vehicle_id=(SELECT vehicle_id FROM users WHERE username=? LIMIT 1)
			ORDER BY created_at ASC`, VIGGO_USERNAME).Scan(&refuels)

		if result.Error != nil || result.RowsAffected == 0 {
			lib.HttpError(w, http.StatusInternalServerError, "Unable to get refuels")
			return
		}
		refuelsCache.Set(VIGGO_USERNAME, refuels, cache.DefaultExpiration)
		err := json.NewEncoder(w).Encode(&refuels)
		if err != nil {
			lib.HttpError(w, http.StatusInternalServerError, "Unable to get refuels")
		}
		return
	}

	err := json.NewEncoder(w).Encode(&cachedRefuels)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to get refuels")
	}
}
