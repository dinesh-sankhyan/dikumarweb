package server

import (
	"net/http"

	"github.com/dikumarweb/logger"
)


func (s *Server) tokenHandler(w http.ResponseWriter, r *http.Request) {

	WriteJSON(w, http.StatusOK, "success")
}