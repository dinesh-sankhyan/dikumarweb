package server

import (

)



//SaveSession save session
func SaveSession(r *http.Request, w http.ResponseWriter, sessName, sessKey string, sessionObj map[interface{}]interface{}) {
	SaveSessionToStore(r, w, sessName, sessKey, sessionTimeOut, sessionObj)
}

//SaveSessionToStore save session
func SaveSessionToStore(r *http.Request, w http.ResponseWriter, sessName, sessKey string, timeout int, sessionObj map[interface{}]interface{}) {
	// Get a session.
	sess, err := Store.Get(r, sessName)
	if err != nil {
		logger.Error(err.Error())
	}
	sess.Values = sessionObj

	if sessionObj != nil {
		logger.Infof("New session created for %s", sessKey)
	}

	sess.Options = &sessions.Options{
		Domain:   domainName,
		Path:     "/",
		MaxAge:   timeout,
		HttpOnly: true,
	}

	if err = sess.Save(r, w); err != nil {
		logger.Fatalf("Error saving session: %v", err)
	}
}
