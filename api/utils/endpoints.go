package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func getEndpoints(router *mux.Router, path string) []byte {
	var routes []string
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		endpoint, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}
		if strings.HasPrefix(endpoint, path) {
			routes = append(routes, endpoint)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return Marshal(routes)
}

func ParseEndpoints(router *mux.Router, path string) {
	router.HandleFunc(path+"/endpoints", func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		res := getEndpoints(router, path)
		_, err := io.Writer.Write(w, res)
		log.Info().Msg(Unmarshal(res))
		if err != nil {
			log.Err(err).Stack().Msg("IO error")
			return
		}
	})
}

func Marshal(data any) []byte {
	v, err := json.Marshal(data)
	if err != nil {
		log.Err(err)
		return nil
	}
	return v
}

func Unmarshal(data []byte) string {
	v, err := strconv.Unquote(string(data))
	if err != nil {
		log.Err(err).Stack().
		return ""
	}
	return v
}
