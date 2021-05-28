package main

import (
	"log"
	"myapp/service"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.Use(service.Middleware)
	router.HandleFunc("/send", service.HandlerSendEmail).Methods(http.MethodPost)

	log.Println("Listen and Serve at port " + port)
	log.Println(http.ListenAndServe(":"+port, router))
}
