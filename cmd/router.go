package main

import (
	"Anxiety/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/registration", handlers.Registration).Methods("POST")
	r.HandleFunc("/auth", handlers.Auth).Methods("POST")
	r.HandleFunc("/logout", handlers.Logout).Methods("GET")

	r.HandleFunc("/services/{category}", handlers.ServiceHandler).Methods("GET")
	r.HandleFunc("/specialists/{id}", handlers.SpecialistsHandler).Methods("GET")
	r.HandleFunc("/fasthelp", handlers.FastHelpHandler).Methods("GET")
	r.HandleFunc("/about", handlers.AboutHandler).Methods("GET")
	r.HandleFunc("/contacts", handlers.ContactsHandler).Methods("GET")

	return r
}
