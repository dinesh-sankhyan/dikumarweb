package server

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"time"

	"github.com/dikumarweb/logger"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/sessions"
	"gopkg.in/boj/redistore.v1"
)

var (
	//Pool Redis pool config
	Pool *redis.Pool
	//Store Redis store
	Store *redistore.RediStore

	domainName         string
	sessionStoreSecret string
	sessionTimeOut     int
)

//InitSessionStore Create session store
func InitSessionStore(redisServer, redisPort, domainName, sessionStoreSecret string) {
	Pool = newPool(redisServer + ":" + redisPort)
	sessionTimeOut = 300

	rediStore, err := redistore.NewRediStoreWithPool(Pool, []byte(sessionStoreSecret))
	rediStore.DefaultMaxAge = sessionTimeOut

	if err != nil {
		panic(err)
	}
	Store = rediStore

	logger.Info("Registering session object")
	gob.Register(&SessionObj{})

	domainName = domainName
	sessionStoreSecret = sessionStoreSecret
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

//CloseSessionstore close session store
func CloseSessionstore() {
	Store.Close()
	Pool.Close()
}

//GetSession get the session
func GetSession(r *http.Request, sessName, sessKey string) (sessObj *SessionObj, err error) {
	// Get a session.
	sess, err := Store.Get(r, sessName)

	if err != nil {
		logger.Error(err.Error())
	}

	if obj, ok := sess.Values[sessKey].(*SessionObj); !ok {
		logger.Errorf("Session not found for key %s" + sessKey)
		err = fmt.Errorf("Session not found for key %s" + sessKey)
	} else {
		sessObj = obj
	}
	return sessObj, err
}

//SaveSession save session
func SaveSession(r *http.Request, w http.ResponseWriter, sessName, sessKey string, sessionObj *SessionObj) {
	SaveSessionToStore(r, w, sessName, sessKey, sessionTimeOut, sessionObj)
}

//SaveSessionToStore save session
func SaveSessionToStore(r *http.Request, w http.ResponseWriter, sessName, sessKey string, timeout int, sessionObj *map[interface{}]interface{}) {
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

