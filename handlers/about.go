package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	about := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Specialist: %v\n", about)
}
