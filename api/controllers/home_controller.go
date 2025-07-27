package controllers

import (
	"net/http"

	"github.com/shekhar-patil/go-blog/api/responses"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.Json(w, http.StatusOK, "Welcome to go-blog")
}
