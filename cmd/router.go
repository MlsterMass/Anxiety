package main

import (
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/services/{category}", ServiceHandler).Methods("GET")
	r.HandleFunc("/specialists/{id}", SpecialistsHandler).Methods("GET")
	r.HandleFunc("/fasthelp/", FastHelpHandler).Methods("GET")
	r.HandleFunc("/about/", AboutHandler).Methods("GET")
	r.HandleFunc("/contacts/", ContactsHandler).Methods("GET")

	return r
}
