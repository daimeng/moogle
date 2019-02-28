package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/daimeng/moogle"
	"github.com/daimeng/ratelimit"
	"github.com/daimeng/ratelimit/clock"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "geocoder"
)

type server struct {
	// db *sql.DB
	clock      *clock.Mock
	dailyLimit ratelimit.Limiter
	secLimit   ratelimit.Limiter
	elmLimit   ratelimit.Limiter
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

func main() {
	port := flag.Int("port", 8080, "Server port")
	step := flag.Bool("step", false, "Manual step clock")

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

	var s server
	if *step {
		clock := clock.NewMock()

		s = server{
			clock:      clock,
			dailyLimit: ratelimit.New(500000, ratelimit.WithClock(clock), ratelimit.WithoutSlack),
			secLimit:   ratelimit.New(100, ratelimit.WithClock(clock)),
			elmLimit:   ratelimit.New(1000, ratelimit.WithClock(clock)),
		}
		http.HandleFunc("/step", s.clockStepHandler)
	} else {
		s = server{
			dailyLimit: ratelimit.New(500000, ratelimit.WithoutSlack),
			secLimit:   ratelimit.New(100),
			elmLimit:   ratelimit.New(1000),
		}
	}

	http.HandleFunc("/maps/api/geocode/json", s.geocodeHandler)
	http.HandleFunc("/maps/api/distancematrix/json", s.distanceMatrixHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))

	// if err != nil && err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }
}
