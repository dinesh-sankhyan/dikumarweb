package server

import (
	"github.com/dikumarweb/services"
	"net/http"

	"github.com/dikumarweb/logger"
)


func (s *Server) tokenHandler(w http.ResponseWriter, r *http.Request) {
	//Code to handle login flow

	//Set session

	token,err := services.CreateUserToken()

	if err != nil {
		logger.Errorf("Error handler")
	}


	WriteJSON(w, http.StatusOK, "success")
}