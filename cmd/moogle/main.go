package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/daimeng/moogle"
	"github.com/daimeng/ratelimit"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "geocoder"
)

type server struct {
	// db *sql.DB
	dailyLimit ratelimit.Limiter
	secLimit   ratelimit.Limiter
	elmLimit   ratelimit.Limiter
}

// TODO: Implement
func (s *server) geocodeHandler(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	res := moogle.GeocodeResponse{}
	json, _ := json.Marshal(res)
	w.Write(json)
}

func parsell(s []string) []moogle.Point {
	l := len(s)
	points := make([]moogle.Point, l)

	for i := 0; i < l; i++ {
		ll := strings.Split(s[i], ",")
		lat, _ := strconv.ParseFloat(ll[0], 64)
		lng, _ := strconv.ParseFloat(ll[1], 64)
		points[i].Lat = float64(lat)
		points[i].Lng = float64(lng)
	}
	return points
}

// func encodell(points []moogle.Point) []string {
// 	s := make([]string, len(points))

// 	for i := 0; i < len(points); i++ {
// 		s[i] = fmt.Sprintf("%f,%f", points[i].Lat, points[i].Lng)
// 	}

// 	return s
// }

func (s *server) distanceMatrixHandler(w http.ResponseWriter, r *http.Request) {
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
		json, _ := json.Marshal(moogle.MATRIX_QUERY_LIMIT)
		w.Write(json)
		return
	}

	q := r.URL.Query()
	originStr := q["origins"]
	destStr := q["destinations"]

	origins := strings.Split(originStr[0], "|")
	dests := strings.Split(destStr[0], "|")

	origP := parsell(origins)
	destP := parsell(dests)

	olen := len(origP)
	dlen := len(destP)

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

	w.Header().Set("Content-Type", "application/json")
	res := moogle.MatrixResponse{
		DestinationAddresses: dests,
		OriginAddresses:      origins,
		Rows:                 rows,
	}
	json, _ := json.Marshal(res)
	w.Write(json)
}

func main() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable",
	// 	host, port, dbname)

	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	s := server{
		dailyLimit: ratelimit.New(500000, ratelimit.WithoutSlack),
		secLimit:   ratelimit.New(100, ratelimit.WithoutSlack),
		elmLimit:   ratelimit.New(1000, ratelimit.WithoutSlack),
	}

	http.HandleFunc("/maps/api/geocode/json", s.geocodeHandler)
	http.HandleFunc("/maps/api/distancematrix/json", s.distanceMatrixHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// if err != nil && err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }
}
