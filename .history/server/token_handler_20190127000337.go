package server

import (
	"net/http"

	"github.com/dikumarweb/logger"
)


func (s *Server) tokenHandler(w http.ResponseWriter, r *http.Request) {
	//Code to handle login flow

	//Set session

	sessionCookie, err := GetSession(r, "session", "sessionID")

	if err != nil {
		logger.Errorf("Error handler")
	}

	if sessionCookie == nil || len(sessionCookie) <= 0 {
		sessobj := make(map[interface{}]interface{})
		sessobj["test"] = 1234
		SaveSession(r, w, "session", "sessionID", sessobj)
	}

	WriteJSON(w, http.StatusOK, "success")
}