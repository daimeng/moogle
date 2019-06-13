package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/daimeng/moogle"
	"github.com/juju/ratelimit"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "geocoder"
)

type server struct {
	// db *sql.DB
	dailyLimit *ratelimit.Bucket
	secLimit   *ratelimit.Bucket
	elmLimit   *ratelimit.Bucket
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

var DAY time.Duration
var SEC time.Duration

func main() {
	DAY, _ := time.ParseDuration("24h")
	SEC, _ := time.ParseDuration("1s")

	port := flag.Int("port", 8080, "Server port")
	// step := flag.Bool("step", false, "Manual step clock")
	day := flag.Int64("daylimit", 500000, "Per day query limit")
	sec := flag.Int64("seclimit", 100, "Per second query limit")
	elm := flag.Int64("elmlimit", 1000, "Per second element limit")

	flag.Parse()

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
		dailyLimit: ratelimit.NewBucketWithQuantum(DAY, *day, *day),
		secLimit:   ratelimit.NewBucketWithQuantum(SEC/10, *sec/10*2, *sec/10),
		elmLimit:   ratelimit.NewBucketWithQuantum(SEC/10, *elm, *elm/10),
	}

	http.HandleFunc("/maps/api/geocode/json", s.geocodeHandler)
	http.HandleFunc("/maps/api/distancematrix/json", s.distanceMatrixHandler)
	// http.HandleFunc("/reset", s.clockResetHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))

	// if err != nil && err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }
}
