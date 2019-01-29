package server

import (
	"net/http"

	"github.com/dikumarweb/logger"
)

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	//Code to handle login flow

	//Set session

	sessionCookie, err := GetSession(r, "session", "sessionID")

	if err != nil {
		logger.Errorf("Error handler")
	}

	if sessionCookie == nil || len(sessionCookie) <0 {
		sessobj := make(map[interface{}]interface{})
		sessobj["test"] = 1234
		SaveSession(r, w, "session", "sessionID", sessobj)
	}

	WriteJSON(w, http.StatusOK, "success")
}

// logout revokes the passed user JWT.
func (s *Server) logout(w http.ResponseWriter, r *http.Request) {

	sessionCookie, err := GetSession(r, "session", "sessionID")
	if err == nil {
		SaveSessionToStore(r, w, "session", "session", 1, sessionCookie)
	} else {
		logger.Printf("logout Error %+v", sessionCookie)
	}

	//logout flow
	//Logout()

}
