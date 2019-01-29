package server

import (
	"net/http"
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

	

		if sessionCookie==nil  {
			sessobj := make(map[interface{}]interface{})
			sessobj["test"] = 1234
			SaveSession(r, w, "session", "sessionID", sessobj)
		}

		WriteJSON(w, http.StatusOK, "success")
}

