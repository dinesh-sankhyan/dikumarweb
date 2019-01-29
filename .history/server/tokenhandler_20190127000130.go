package server

import (

)


//SaveSession save session
func SaveSession(r *http.Request, w http.ResponseWriter) {
	SaveSessionToStore(r, w, sessName, sessKey, sessionTimeOut, sessionObj)
}