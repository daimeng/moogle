package main

import (
	"encoding/json"
	"net/http"

	"github.com/daimeng/moogle"
)

// TODO: Implement
func (s *server) geocodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sec, _ := s.secLimit.TryTake()
	elm := false
	if sec {
		elm, _ = s.elmLimit.TryTake()
	}
	daily := false
	if elm {
		daily, _ = s.dailyLimit.TryTake()
	}

	if !(sec && daily && elm) {
		json, _ := json.Marshal(moogle.GEOCODE_QUERY_LIMIT)
		w.Write(json)
		return
	}

	res := moogle.GeocodeResponse{}
	json, _ := json.Marshal(res)
	w.Write(json)
}
