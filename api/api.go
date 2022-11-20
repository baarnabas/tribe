package vaultapi

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"tribe/api/utils"
)

var router = mux.NewRouter()

func CreateApi() {
	apiPort, exists := os.LookupEnv("REST_API_PORT")
	if exists {
		log.Info().Msg("Creating REST API")
		router := router

		/* Initialize package endpoints */
		Healthz(router)
		utils.ParseEndpoints(router, "/")

		handler := cors.New(cors.Options{
			AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "UPDATE", "OPTIONS"},
		}).Handler(router)

		err := http.ListenAndServe(":"+apiPort, handler)

		if err != nil {
			log.Err(err).Stack().Msg("Error occurred creating http server")
			return
		}
	} else {
		log.Info().Msg("Skipping REST API")
	}

}
