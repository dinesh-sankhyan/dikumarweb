package server

import (

)


//tokenHandler save session
func (s *Server) tokenHandler(r *http.Request, w http.ResponseWriter) {
	SaveSessionToStore(r, w, sessName, sessKey, sessionTimeOut, sessionObj)
}