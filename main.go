package main

import (
	"github.com/rs/zerolog/log"
	vaultapi "tribe/api"
)

func main() {
	initApis()
}

func initApis() {
	log.Info().Msg("Initializing Rest API")
	vaultapi.CreateApi()
}
