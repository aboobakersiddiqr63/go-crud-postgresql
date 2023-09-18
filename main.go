package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/aboobakersiddiqr63/go-crud-postgresql/routes"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	fmt.Println("Go API using PostgreSQL")

	app := router.Router()
	fmt.Println("Starting the server on port 4000")

	log.Fatal(http.ListenAndServe(":4000", app))
}
