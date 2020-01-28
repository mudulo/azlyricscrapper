package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"encoding/json"
	"log"
	"fmt"
	"golyrics/controller"
)

func search(w http.ResponseWriter, r *http.Request){

	query := r.URL.Query().Get("song")
	fmt.Println(query)
	response, err := json.Marshal(controller.SearchLyrics(query))
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}

	w.Header().Set("Content-Type", "application/json")
  	w.Write(response)
}

func main() {
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
	router := mux.NewRouter()
	router.HandleFunc("/search", search).Methods("GET")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))

}