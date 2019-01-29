package server

import (
	"net/http"
)

type LoginParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {

	params := &LoginParams{}
	//Code to handle login flow

	//Set session

	sessionCookie, err := services.GetSession(r, "session", resp.SessionXID)
		if err != nil {
			logger.Errorf("login error while setting cookie for userid %s, err %+v", params.Username, err)
		}

		if sessionCookie==nil || len(sessionCookie) == 0 {
			sessobj := make(map[interface{}]interface{})
			sessobj["accesstoken"] = resp.AccessToken
			sessobj["refreshtoken"] = resp.RefreshToken
			sessobj["sessionXID"] = resp.SessionXID
			sessobj["test"] = 1234
			services.SaveSession(r, w, "session", resp.SessionXID, sessobj)
		}

		WriteJSON(w, http.StatusOK, resp)

}

