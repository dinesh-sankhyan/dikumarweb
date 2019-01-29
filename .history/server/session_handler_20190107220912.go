package server

import (
	"net/http"

	"github.com/dikumarweb/logger"
)

type LoginParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {

	//params := &LoginParams{}
	//Code to handle login flow

	//Set session

	sessionCookie, err := GetSession(r, "session", "sessionID")

	if err != nil {
		logger.Errorf("Error handler")
	}

	if sessionCookie == nil {
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

	err = services.Logout(userJWT)

}