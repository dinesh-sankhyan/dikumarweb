package server

import (
	"net/http"

	"github.com/dikumarweb/logger"
)


func (s *Server) tokenHandler(w http.ResponseWriter, r *http.Request) {
	//Code to handle login flow

	//Set session

	services.

	if err != nil {
		logger.Errorf("Error handler")
	}


	WriteJSON(w, http.StatusOK, "success")
}