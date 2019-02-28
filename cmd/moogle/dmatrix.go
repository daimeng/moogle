package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/daimeng/moogle"
)

func (s *server) distanceMatrixHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	key := q.Get("key")
	if key == "" {
		json, _ := json.Marshal(moogle.MATRIX_NO_KEY)
		w.Write(json)
		return
	}

	originStr := q.Get("origins")
	destStr := q.Get("destinations")

	origins := strings.Split(originStr, "|")
	dests := strings.Split(destStr, "|")

	olen := len(origins)
	dlen := len(dests)

	ready, _ := s.secLimit.TryTake(1)
	if !ready {
		json, _ := json.Marshal(moogle.MATRIX_QUERY_LIMIT)
		w.Write(json)
		return
	}

	ready, _ = s.elmLimit.TryTake(olen * dlen)
	if !ready {
		json, _ := json.Marshal(moogle.MATRIX_ELEMENT_LIMIT)
		w.Write(json)
		return
	}

	ready, _ = s.dailyLimit.TryTake(1)
	if !ready {
		json, _ := json.Marshal(moogle.MATRIX_DAILY_LIMIT)
		w.Write(json)
		return
	}

	origP := parsell(origins)
	destP := parsell(dests)

	dists := moogle.DistManhattan(origP, destP)
	rows := make([]moogle.DistanceRow, olen)

	for i := range origins {
		elems := make([]moogle.DistanceElm, dlen)
		for j := range dests {
			d := int(dists[i*dlen+j])
			elems[j].Distance = &moogle.TextedInt{
				Value: d,
				Text:  fmt.Sprintf("%d m", d),
			}
			elems[j].Duration = &moogle.TextedInt{
				Value: d / 20,
				Text:  fmt.Sprintf("%d min", d/20),
			}
			elems[j].Status = moogle.ElmOk
		}
		rows[i].Elements = elems
	}

	res := moogle.MatrixResponse{
		DestinationAddresses: dests,
		OriginAddresses:      origins,
		Rows:                 rows,
	}
	json, _ := json.Marshal(res)
	log.Printf("%s", json)
	w.Write(json)
}
