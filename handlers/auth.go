package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User : ", r.Form.Get("name"))

}

func Auth(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Specialist: %v\n", user)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Specialist: %v\n", user)

}
