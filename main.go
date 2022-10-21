package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		log.Err(err).Stack().Msg("IO error")
		return
	}
	log.Info().Msg("Health is OK")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "UPDATE", "OPTIONS"},
	}).Handler(router)

	err := http.ListenAndServe(":8000", handler)

	if err != nil {
		log.Err(err).Stack().Msg("Error occurred creating http server")
		return
	}
}
