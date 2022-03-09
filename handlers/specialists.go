package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func SpecialistsHandler(w http.ResponseWriter, r *http.Request) {
	specialist := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Specialist: %v\n", specialist["id"])

}
