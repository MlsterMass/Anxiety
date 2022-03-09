package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	contact := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Specialist: %v\n", contact)
}
