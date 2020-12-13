package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/maritza05/golang-bootcamp-2020/satellites/adapters/repos"
	"github.com/maritza05/golang-bootcamp-2020/satellites/adapters/rest"
	"github.com/maritza05/golang-bootcamp-2020/satellites/app"
)

func RegisterListSatellitesRoute(router *mux.Router, path string) {
	csvFile, err := os.Open("satellites.csv")
	if err != nil {
		panic("Can't open satellites.csv file")
	}
	repository, err := repos.NewCsvRepo(csvFile)
	if err != nil {
		panic("Error while processing csv file")
	}
	service := app.NewListSatelliteService(repository)
	controller := rest.NewListSatelliteController(service)
	router.HandleFunc(path, controller.GetAll).Methods(http.MethodGet)
}
