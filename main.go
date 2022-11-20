package main

import (
	"github.com/rs/zerolog"
	"os"
	vaultapi "tribe/api"
)

func main() {
	setLogging()
	initRestApi()
}

func setLogging() {
	level, exists := os.LookupEnv("LOGGING_LEVEL")
	if exists {
		switch level {
		case "trace":
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warn":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		}
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func initRestApi() {
	state, exists := os.LookupEnv("REST_API")
	if exists && state == "enabled" {
		vaultapi.CreateApi()
	}
}
