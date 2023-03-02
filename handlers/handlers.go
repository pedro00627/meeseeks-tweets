package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers(){
	router := mux.NewRouter()

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+Port, handler))

}