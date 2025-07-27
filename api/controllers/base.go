package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/shekhar-patil/go-blog/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize(dbHost, dbUser, dbPassword, dbName, dbPort string) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	fmt.Println("DSN:", dsn)
	var err error
	s.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to postgres database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the postgres database")
	}

	s.DB.AutoMigrate(&models.User{}, &models.Post{})
	s.Router = mux.NewRouter()
	s.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
