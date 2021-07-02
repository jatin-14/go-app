package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jatin-14/go-app/apis"
)

func main(){

	router := mux.NewRouter()
	router.HandleFunc("/movies", apis.GetMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", apis.GetMovie).Methods("GET")
	router.HandleFunc("/movie", apis.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", apis.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movie/{id}", apis.UpdateMovie).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("server running")

}