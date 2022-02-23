package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func handleIndexRoute(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello Kubernetes")
}

func main() {
	godotenv.Load()

	PORT, _ := strconv.Atoi(os.Getenv("PORT"))

	URL := fmt.Sprintf("0.0.0.0:%d", PORT)

	r := mux.NewRouter()

	r.HandleFunc("/", handleIndexRoute)

	srv := &http.Server{
		Addr:         URL,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.Printf("Listening for requests on port :%d", PORT)
	log.Fatal(srv.ListenAndServe())
}
