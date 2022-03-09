package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	router := InitRouter()
	fmt.Printf("Starting server at port 8888\n")
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
