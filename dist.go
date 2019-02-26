package moogle

import (
	"math"
)

const circ = 2 * math.Pi * 6371000

func DistManhattan(origs []Point, dests []Point) []float64 {
	dlen := len(dests)
	olen := len(origs)
	dists := make([]float64, olen*dlen)

	for i, o := range origs {
		for j, d := range dests {
			dists[i*dlen+j] = (math.Abs(o.Lat-d.Lat) + math.Abs(o.Lng-d.Lng)) / 180 * circ
		}
	}
	return dists
}
