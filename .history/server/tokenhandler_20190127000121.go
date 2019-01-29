package server

import (

)


//SaveSession save session
func SaveSession(r *http.Request, w http.ResponseWriter, sessName, sessKey string, sessionObj map[interface{}]interface{}) {
	SaveSessionToStore(r, w, sessName, sessKey, sessionTimeOut, sessionObj)
}