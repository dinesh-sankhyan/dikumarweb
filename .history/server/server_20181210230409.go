// Package server provides the web server implementation
package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"context"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// Server wires up routes and starts api server
type Server struct {
	swagger []byte
	http.Server
	shutdownReq chan bool
	reqCount    uint32
}

// New provides a new Server
func New() *Server {
	s := &Server{
		Server: http.Server{
			Addr:         ":8080",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  10 * time.Second,
		},
		shutdownReq: make(chan bool),
	}
	router := s.RegisterHandlers
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*.mheducation.com", "*localhost:8080"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	s.Handler = handlers.CORS(headersOk, originsOk, methodsOk)(router())
	return s
}

// RegisterHandlers wires up handlers
func (s *Server) RegisterHandlers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Handlers for static content
	router.PathPrefix("/static/").Handler(s.staticHandler("/static/"))
	router.PathPrefix("/swagger-ui/").Handler(s.staticHandler("/swagger-ui/"))

	// Define the handlers for application paths
	for _, route := range s.routes() {
		router.Handle(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
	s.Handler = router

	//Attach profiler

	AttachProfiler(router)

	return router
}

//AttachProfiler attach profiler routes
func AttachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}

//WaitShutdown wait for server shutdown
func (s *Server) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		fmt.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-s.shutdownReq:
		fmt.Printf("Shutdown request (/shutdown %v)", sig)
	}

	//logger.Infof("Stopping http server ...")

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		//logger.Infof("Shutdown request error: %v", err)
	}
}

//Shutdown shutdown server call
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}
