package vaultapi

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"tribe/api/healthz"
)

func CreateApi() {
	value, exists := os.LookupEnv("VAULT_API_PORT")
	if exists {
		log.Info().Msg("Creating Vault API")
		router := mux.NewRouter()

		/* Initialize package endpoints */
		healthz.Init(router)

		handler := cors.New(cors.Options{
			AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "UPDATE", "OPTIONS"},
		}).Handler(router)
		err := http.ListenAndServe(":"+value, handler)

		if err != nil {
			log.Err(err).Stack().Msg("Error occurred creating http server")
			return
		}
	} else {
		log.Info().Msg("Skipping Vault API")
	}

}
