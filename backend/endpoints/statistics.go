package endpoints

import (
	"encoding/json"
	"net/http"
	"services/project/lib"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type FrequencyStats struct {
	Year         string      `json:"year"`
	TotalCost    float32     `json:"total_cost"`
	TotalTrip    float32     `json:"total_trip"`
	TotalRefuels float32     `json:"total_refuels"`
	Weekly       []ShortStat `json:"weekly"`
	Monthly      []ShortStat `json:"monthly"`
}

type ShortFrequencyStats struct {
	Year         string  `json:"year"`
	TotalCost    float32 `json:"total_cost"`
	TotalTrip    float32 `json:"total_trip"`
	TotalRefuels float32 `json:"total_refuels"`
}

type ShortStat struct {
	DateValue int     `json:"date_value"`
	SumCost   float32 `json:"sum_cost"`
	SumTrip   float32 `json:"sum_trip"`
	Count     int     `json:"count"`
}

// @TODO
func GetRefuelFrequencyForPeriod(w http.ResponseWriter, r *http.Request) {
	year := r.URL.Query().Get("year")
	if year == "" {
		year = strconv.Itoa(time.Now().Year())
	}
	userId := chi.URLParam(r, "userId")

	var stat FrequencyStats
	var shortStat ShortFrequencyStats

	db := lib.GetDatabase()

	result := db.Raw(`
		SELECT strftime('%Y', r.fuel_date) AS year, SUM(r.cost) AS total_cost, SUM(r.trip_km) AS total_trip, COUNT(r.cost) AS total_refuels
		FROM users u
		INNER JOIN vehicles v	on u.id = v.user_id
		INNER JOIN refuels r on v.id = r.vehicle_id
		WHERE u.id = ?
		AND year = ?`, userId, year).Scan(&shortStat)

	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to get statistics")
		return
	}

	stat.Year = shortStat.Year
	stat.TotalCost = shortStat.TotalCost
	stat.TotalTrip = shortStat.TotalTrip
	stat.TotalRefuels = shortStat.TotalRefuels

	var weekly []ShortStat
	result = db.Raw(`
		SELECT strftime('%Y', r.fuel_date) AS year, strftime('%W', r.fuel_date) AS date_value, SUM(r.cost) AS sum_cost, SUM(r.trip_km) AS sum_trip, COUNT(r.cost) AS count
		FROM users u
		INNER JOIN vehicles v	on u.id = v.user_id
		INNER JOIN refuels r on v.id = r.vehicle_id
		WHERE u.id = ?
		AND year = ?
		GROUP BY date_value
		ORDER BY date_value ASC`, userId, year).Scan(&weekly)
	stat.Weekly = weekly

	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to get statistics")
		return
	}

	var monthly []ShortStat
	result = db.Raw(`
		SELECT strftime('%Y', r.fuel_date) AS year, strftime('%m', r.fuel_date) AS date_value, SUM(r.cost) AS sum_cost, SUM(r.trip_km) AS sum_trip, COUNT(r.cost) AS count
		FROM users u
		INNER JOIN vehicles v	on u.id = v.user_id
		INNER JOIN refuels r on v.id = r.vehicle_id
		WHERE user_id = ?
		AND year = ?
		GROUP BY date_value
		ORDER BY date_value ASC`, userId, year).Scan(&monthly)
	stat.Monthly = monthly

	if result.Error != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to get statistics")
		return
	}

	err := json.NewEncoder(w).Encode(&stat)
	if err != nil {
		lib.HttpError(w, http.StatusInternalServerError, "Unable to get statistics")
	}
}

func TrackFirstVisit(w http.ResponseWriter, r *http.Request) {

}
