package vaultapi

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"tribe/api/utils"
)

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		log.Err(err).Stack().Msg("IO error")
		return
	}
	log.Info().Msg("Health is OK")
}

func Healthz(router *mux.Router) {
	router.HandleFunc("/healthz", healthz).Methods("GET")
	utils.ParseEndpoints(router, "/healthz")
}
