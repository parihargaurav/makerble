package routes

import (
    "github.com/gorilla/mux"
    "myapp/controllers"
)

func RegisterRoutes() *mux.Router {
    router := mux.NewRouter()

    // Define your routes and map them to controller functions
    router.HandleFunc("/login", controllers.Login).Methods("POST")
    router.HandleFunc("/register", controllers.Register).Methods("POST") // Registration route
    router.HandleFunc("/patients", controllers.RegisterPatient).Methods("POST")
    router.HandleFunc("/patients/{id}", controllers.GetPatient).Methods("GET")
	router.HandleFunc("/patients/{id}", controllers.UpdatePatient).Methods("PUT") // Added PUT route
	router.HandleFunc("/patients/{id}", controllers.DeletePatient).Methods("DELETE") // Add DELETE route

    // Other routes can be added here

    return router
}
