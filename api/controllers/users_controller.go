package controllers

import (
	"net/http"

	"github.com/shekhar-patil/go-blog/api/models"
	"github.com/shekhar-patil/go-blog/api/responses"
)

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	users, err := user.FindAllUsers(s.DB)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, users)
}
