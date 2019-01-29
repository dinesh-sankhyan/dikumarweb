package server

import (
	"os"
	"net/http"
)

func (s *Server) dummyRoute(w http.ResponseWriter, r *http.Request) {

	// do some additional test of health here. For now, respond 200
	health := Health{}
	health.Body.Ok = true
	health.Body.Messages = []HealthMessage{{"application", "OK"}}
	host, _ := os.Hostname()
	health.Body.HostName = host
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, api_key, Authorization")

	if r.Method == "HEAD" {
		w.WriteHeader(http.StatusOK)
	} else {
		WriteJSON(w, http.StatusOK, health.Body)
	}
}

func S