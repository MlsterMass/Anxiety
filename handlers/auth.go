package handlers

import (
	"Anxiety/repository"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func RegistrationUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	children, err := strconv.ParseBool(r.FormValue("children"))
	if err != nil {
		fmt.Println("Incorrect value")
	}
	pets, err := strconv.ParseBool(r.FormValue("pets"))
	if err != nil {
		fmt.Println("Incorrect value")
	}
	user := repository.Users{
		Name:      r.FormValue("name"),
		Nickname:  r.FormValue("nickname"),
		Gender:    r.FormValue("gender"),
		Status:    r.FormValue("status"),
		Childrens: children,
		Pets:      pets,
		Location:  r.FormValue("location"),
		Password:  r.FormValue("password"),
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(user)
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
