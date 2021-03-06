package server

import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dikumarweb/logger"
	"github.com/dikumarweb/services"
)

// middleware is a definition of  what a middleware is,
// take in one handlerfunc and wrap it within another handlerfunc
type middleware func(http.HandlerFunc) http.HandlerFunc

// BuildChain builds the middlware chain recursively, functions are first class
func BuildChain(f http.HandlerFunc, m ...middleware) http.HandlerFunc {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise nest the handlerfuncs
	return m[0](BuildChain(f, m[1:cap(m)]...))
}

// MiddleWareChain PublicChain create an endpoint grouping called publicChain
// which has the public middlewares
var MiddleWareChain = []middleware{
	FuncExecTimeMiddleware,
	CORSMiddleware,
	AuthMiddleware,
}

// MiddleWareNoAuthChain PublicChain create an endpoint grouping called publicChain
// which has the public middlewares
var MiddleWareNoAuthChain = []middleware{
	FuncExecTimeMiddleware,
	CORSMiddleware,
}

// AuthMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
//This function is a wrapper to all the request to validate jwt token
var AuthMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// pre handler functionality
		urlInfo := methodURIString(r)
		logger.Println("AuthMiddleware : Start auth for " + urlInfo)
		isValidToken, _ := services.JwtValidateToken(r)
		if !isValidToken {
			WriteJSON(w, http.StatusUnauthorized, http.StatusUnauthorized)
			logger.Println("AuthMiddleware: End auth  for " + urlInfo)
			return
		}
		f(w, r)
		logger.Println("AuthMiddleware: End auth  for " + urlInfo)
		// post handler functionality
	}
}

// FuncExecTimeMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
//This function is a wrapper to log request execution time
var FuncExecTimeMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// ... pre handler functionality
		startTime := time.Now()
		urlInfo := methodURIString(r)
		logger.Info("FuncExecTimeMiddleware : Start Handling request for " + urlInfo)
		f(w, r)
		logger.Infof("FuncExecTimeMiddleware: Total request Exec time for " +
			urlInfo + time.Since(startTime).String())
		// ... post handler functionality
	}
}

// CORSMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
//This function is a filter for cross irgin requests
var CORSMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// ... pre handler functionality
		logger.Info("CORSMiddleware : start")
		origin := r.Header.Get("Origin")
		name, _ := os.Hostname()
		w.Header().Set("Host", name)
		if isAllowedDomain(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, "+
				"Accept-Encoding, Accept-Language, X-CSRF-Token, Authorization")
		}
		if r.Method == "OPTIONS" {
			logger.Infof("CORSMiddleware: Done")
			return
		}
		f(w, r)
		logger.Infof("CORSMiddleware: Done")
		// ... post handler functionality
	}
}

func methodURIString(r *http.Request) string {
	var buffer bytes.Buffer
	buffer.WriteString("Method '")
	buffer.WriteString(r.Method)
	buffer.WriteString("' and uri '")
	buffer.WriteString(r.URL.Path)
	buffer.WriteString("' ")

	return buffer.String()
}
func isAllowedDomain(origin string) bool {
	domains := "localhost,localhost:4200,.xyz.com"

	allowedDomains := strings.Split(domains, ",")

	if origin != "" {
		for _, v := range allowedDomains {
			if strings.Contains(origin, v) {
				return true
			}
		}
	}
	return false
}
