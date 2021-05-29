package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores es usado para conectarse*/
func Manejadores() {
	fmt.Println("Es el nuevo")
	router := mux.NewRouter()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handle(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
