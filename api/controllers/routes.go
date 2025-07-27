package controllers

import "github.com/shekhar-patil/go-blog/api/middlewares"

func (s *Server) initializeRoutes()	{
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
}