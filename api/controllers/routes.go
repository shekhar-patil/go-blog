package controllers

import (
	"github.com/shekhar-patil/go-blog/api/middlewares"
)

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("Get")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.UpdateUser)).Methods("PUT")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE")
}
