package endpoints

import (
	"fmt"
	"net/http"
)

// @TODO
func GetRefuelFrequencyForPeriod(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	span := r.URL.Query().Get("span")

	fmt.Println("statistics" + from + to + span)

	/*
	sqlite> select strftime('%W', created_at) AS week, SUM(trip_km) from refuels group by week ORDER BY week ASC;
	*/
}

func TrackFirstVisit(w http.ResponseWriter, r *http.Request) {

}
