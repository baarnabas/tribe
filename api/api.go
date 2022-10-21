package vaultapi

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
	"strings"
	"tribe/api/healthz"
)

var router = mux.NewRouter()

func getEndpoints(router *mux.Router) []string {
	var routes []string
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}
		routes = append(routes, path)
		return nil
	})
	if err != nil {
		return nil
	}
	return routes
}
func healthzEndpoints(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	_, err := io.WriteString(w, "[\""+strings.Join(getEndpoints(router), "\",\"")+"\"]")
	if err != nil {
		log.Err(err).Stack().Msg("IO error")
		return
	}
	log.Info().Msg("Health is OK status")
}

func CreateApi() {
	value, exists := os.LookupEnv("VAULT_API_PORT")
	if exists {
		log.Info().Msg("Creating Vault API")
		router := router

		/* Initialize package endpoints */
		healthz.Init(router)
		router.HandleFunc("/healthz/endpoints", healthzEndpoints).Methods("GET")

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
