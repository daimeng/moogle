package main

import (
	"encoding/json"
	"net/http"

	"github.com/daimeng/moogle"
)

// TODO: Implement
func (s *server) geocodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, ready := s.secLimit.TakeMaxDuration(1, 10*SEC)
	if !ready {
		json, _ := json.Marshal(moogle.GEOCODE_QUERY_LIMIT)
		w.Write(json)
		return
	}

	// ready, _ = s.dailyLimit.TryTake(1)
	// if !ready {
	// 	json, _ := json.Marshal(moogle.GEOCODE_DAILY_LIMIT)
	// 	w.Write(json)
	// 	return
	// }

	res := moogle.GeocodeResponse{}
	json, _ := json.Marshal(res)
	w.Write(json)
}
