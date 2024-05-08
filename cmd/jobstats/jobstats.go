package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	_, err := initDB(os.Getenv("DB_URL"), os.Getenv("DB_LOCAL") == "true")
	if err != nil {
		log.Fatal(err)
	}

	routerV1 := http.NewServeMux()
	routerV1.HandleFunc("/ready", handleWelcome)
	routerV1.HandleFunc("/error", handleError)
	routerV1.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		respondError(w, http.StatusNotFound, "not found")
	})

	router := http.NewServeMux()
	router.Handle("/v1/", http.StripPrefix("/v1", routerV1))
	router.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		respondError(w, http.StatusNotFound, "not found")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	welcome := struct {
		Message string `json:"message"`
	}{
		Message: "welcome to the jobstats API",
	}
	respondJSON(w, http.StatusOK, welcome)
}

func handleError(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusTeapot, "server gives error as expected")
}

func respondJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, statusCode int, errMsg string) {
	type err struct {
		Error string `json:"error"`
	}
	respondJSON(w, statusCode, err{Error: errMsg})
}
