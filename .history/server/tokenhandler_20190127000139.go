package server

import (

)


//SaveSession save session
func tokenHandler(r *http.Request, w http.ResponseWriter) {
	SaveSessionToStore(r, w, sessName, sessKey, sessionTimeOut, sessionObj)
}