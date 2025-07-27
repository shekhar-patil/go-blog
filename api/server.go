package api
import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"github.com/shekhar-patil/go-blog/api/controllers"
)

var server = controllers.Server{}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	server.Initialize(dbHost, dbUser, dbPassword, dbName, dbPort)
	server.Run(":8080")
}