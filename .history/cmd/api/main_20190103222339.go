/*
Package main Module API

Your Project description goes here.

Security:
	- api_key:

SecurityDefinitions:
    api_key:
        type: apiKey
        name: Authorization
        in: header
	Version:   1.0.0
	Host:      @@HOSTPORT@@
	License: Proprietary API http://www.yoursite.com/terms-use.html

Consumes:
	- application/json

Produces:
	- application/json

swagger:meta
*/
package main

import (
	"github.com/dikumarweb/logger"
	"fmt"

	"github.com/dikumarweb/server"
)

//go:generate swagger generate spec -i ../../tags.json -o ../swagger.json
func main() {
	logger.InitLogger("debug", "appName", "Version1.o", "file.txt", "")
	serverInitSessionStore("localhost", "6379", "localhost:8080", "sessionSecret")

	logger.Info("Starting...")
	s := server.New()
	done := make(chan bool)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			fmt.Printf("Listen and serve: %v", err)
		}
		done <- true
	}()

	//wait shutdown
	s.WaitShutdown()

	<-done
}
