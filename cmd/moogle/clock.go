package main

import (
	"net/http"
	"time"
)

func (s *server) clockStepHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	t, _ := time.ParseDuration(q.Get("t"))
	s.clock.Move(t)
}