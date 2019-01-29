package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dikumarweb/logger"
	"github.com/gorilla/mux"
)

// Route holds route metadata
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes getter
func (s *Server) routes() []Route {
	var routes = []Route{
		// Route for root page
		{"Root", "GET", "/", s.rootHandler},
		{"Root", "GET", "/swagger-ui/", s.rootHandler},
		{"SwaggerJson", "GET", "/swagger.json", s.swaggerJSONHandler},

		//ShutDownHandler
		//{"Root", "GET", "/shutdown", s.ShutdownHandler},

		// Route for health requests
		{"HealthShow", "GET", "/health", s.healthHandler},
		{"HealthShow", "HEAD", "/health", s.healthHandler},
		{"Dummy", "GET", "/dummy", BuildChain(s.dummyRouteHandler, MiddleWareChain...)},
		{"Dummy", "GET", "/calc/{operation}/", BuildChain(s.calcRouteHandler, MiddleWareChain...)},
	}

	return routes
}

// rootHandler serves the root page
func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/swagger-ui/")
	WriteJSON(w, http.StatusMovedPermanently, "Forwarding to Swagger Documentation")
}

// healthHandler returns service health
// swagger:route GET /health Health GetHealth
//
// Health Check
//
// This will return the health of the service and its subsystems.
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: HealthSuccess
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {

	// do some additional test of health here. For now, respond 200
	health := Health{}
	health.Body.Ok = true
	health.Body.Messages = []HealthMessage{{"application", "OK"}}
	host, _ := os.Hostname()
	health.Body.HostName = host
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, api_key, Authorization")

	if r.Method == "HEAD" {
		w.WriteHeader(http.StatusOK)
	} else {
		WriteJSON(w, http.StatusOK, health.Body)
	}
}

//Health health end point
// swagger:response HealthSuccess
type Health struct {
	// in:body
	Body struct {
		// Is the service healthy?
		Ok bool `json:"ok"`
		// Significant messages from subsystems
		Messages []HealthMessage `json:"messages"`
		// SCM commit ID for the currently running build
		HostName string `json:"hostname"`
	}
}

// HealthMessage from a subsystem
type HealthMessage struct {
	// Name of the subsystem
	Subsystem string `json:"subsystem"`
	// Message from the subsystem
	Message string `json:"message"`
}

const staticFilePath = "/Users/dinesh_kumar/go/src/github.com/dikumarweb/swagger/static/"

// staticHandler creates a handler for static content
func (s *Server) staticHandler(path string) http.Handler {
	fs := http.FileServer(http.Dir(staticFilePath))
	return http.StripPrefix(path, fs)
}

// swaggerJSONHandler serves the swagger.json file (generated)
func (s *Server) swaggerJSONHandler(w http.ResponseWriter, r *http.Request) {
	if s.swagger == nil {
		// Read the file into memory and substitute the current host and port
		rawSwagger, err := ioutil.ReadFile(staticFilePath + "swagger.json")
		if err != nil {
			logger.WithError(err).Error("Unable to read swagger.json")
		}
		s.swagger = bytes.Replace(rawSwagger, []byte("@@HOSTPORT@@"), []byte(r.Host), -1)
	}
	http.ServeContent(w, r, "swagger.json", time.Unix(0, 0), bytes.NewReader(s.swagger))
}

func WriteJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	if v == nil {
		w.WriteHeader(statusCode)
	} else {
		w.Header().Set("Content-Type", "application/json")
		b, err := json.Marshal(v)
		if err != nil {
			WriteError(w, http.StatusInternalServerError, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		_, err = w.Write(b)
		//Handle Error.
		if err != nil {
			//logger.Errorf("WriteJSON Error while writing JSON %s, The Object string is %s", err,string(b))
		}
	}
}

// WriteError forms an ErrorResponse and writes it as a response.
func WriteError(w http.ResponseWriter, statusCode int, err error) {
	e := ErrorResponse{}
	e.Body.Error = err.Error()
	e.Body.Code = statusCode
	WriteJSON(w, statusCode, e.Body)
}

// ErrorResponse error response
//swagger:response errorResponse
type ErrorResponse struct {
	// in:body
	Body struct {
		// The error message
		Error string `json:"error"`
		// HTTP status code, same as on response
		Code int `json:"code"`
	}
}

// dummyRouteHandler returns json response
// swagger:route GET /dummy Health GetHealth
//
// Health Check
//
// This will return the health of the service and its subsystems.
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: HealthSuccess
func (s *Server) dummyRouteHandler(w http.ResponseWriter, r *http.Request) {

	// do some additional test of health here. For now, respond 200
	health := Health{}
	health.Body.Ok = true
	health.Body.Messages = []HealthMessage{{"application", "OK"}}
	host, _ := os.Hostname()
	health.Body.HostName = host
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, api_key, Authorization")

	if r.Method == "HEAD" {
		w.WriteHeader(http.StatusOK)
	} else {
		WriteJSON(w, http.StatusOK, health.Body)
	}
}

// calcRouteHandler returns result of operation
// swagger:route GET /calc/{operation} Health GetHealth
//
// Calculato service
//
// This will return the result of operation on 2 numbers.
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: HealthSuccess
func (s *Server) calcRouteHandler(w http.ResponseWriter, r *http.Request) {
	var op1, op2 float64
	operation := "add"

	oper := mux.Vars(r)["operation"]
	if oper != "" {
		operation= oper
	}

	param2 := mux.Vars(r)["op2"]
	if param2 != "" {
		op2, _ = strconv.ParseFloat(param2, 64)
	}

	varsQ := r.URL.Query()
	
	if val, ok := varsQ["operation"]; ok {
		operation = val[0]
	}

	calc := CalcService{}

	var result float64
	if operation == "add" {
		result = calc.AddValue(op1, op2)
	} else if operation == "multiply" {
		result = calc.MultiplyValue(op1, op2)
	}

	WriteJSON(w, http.StatusOK, result)
}
