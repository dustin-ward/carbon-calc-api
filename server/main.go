// Containerized Golang API service
package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/APIDemo/database"
	"example.com/APIDemo/routes"
)

func main() {

	database.Connect()
	fmt.Println("Database Connection Established")
	defer database.Disconnect()

	fmt.Println("API Service Running")

	router := routes.Setup()

	log.Fatal(http.ListenAndServe(":8080", router))
}
