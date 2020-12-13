package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Start() {
	loadEnvVariables()
	router := mux.NewRouter()

	RegisterListSatellitesRoute(router, "/satellites")
	RegisterExportSatellitesRoute(router, "/export")

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func loadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while looking for .env file")
	}
}
