package healthz

import (
	"github.com/gorilla/mux"
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

func Init(router *mux.Router) {
	router.HandleFunc("/healthz", Healthz).Methods("GET")
}
