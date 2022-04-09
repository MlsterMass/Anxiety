package main

import (
	"Anxiety/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Registration
	r := router.PathPrefix("/registration").Subrouter()
	r.HandleFunc("/user", handlers.RegistrationUser).Methods("POST")
	r.HandleFunc("/volunteer", handlers.RegistrationVolunteer).Methods("POST")
	r.HandleFunc("/specialist", handlers.RegistrationSpecialist).Methods("POST")
	r.HandleFunc("/support", handlers.RegistrationSupport).Methods("POST")

	// Login
	l := router.PathPrefix("/login").Subrouter()
	l.HandleFunc("/user", handlers.LoginUser).Methods("POST")
	l.HandleFunc("/volunteer", handlers.LoginVolunteer).Methods("POST")
	l.HandleFunc("/specialist", handlers.LoginSpecialist).Methods("POST")
	l.HandleFunc("/support", handlers.LoginSupport).Methods("POST")

	router.HandleFunc("/logout", handlers.Logout).Methods("GET")
	//
	//r.HandleFunc("/services/{category}", handlers.ServiceHandler).Methods("GET")
	//r.HandleFunc("/specialists/{id}", handlers.SpecialistsHandler).Methods("GET")
	//r.HandleFunc("/fasthelp", handlers.FastHelpHandler).Methods("GET")
	//r.HandleFunc("/about", handlers.AboutHandler).Methods("GET")
	//r.HandleFunc("/contacts", handlers.ContactsHandler).Methods("GET")

	return router
}
