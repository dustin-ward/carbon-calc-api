package routes

import (
	"example.com/APIDemo/controllers"
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.HomePage)

	// Auth routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	// Learning routes
	router.HandleFunc("/calc", controllers.CarbonCalculator).Methods("POST")
	router.HandleFunc("/evaluate", controllers.EvaluateChain).Methods("POST")

	// Blockchain routes
	router.HandleFunc("/temp", controllers.Temp)

	return router
}
