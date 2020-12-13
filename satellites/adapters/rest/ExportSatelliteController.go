package rest

import (
	"net/http"
	"os"

	"github.com/maritza05/golang-bootcamp-2020/satellites/app/ports/in"
)

type FileHandler interface {
	CreateFile(filename string) *os.File
	Close()
}

type ExportSatelliteController struct {
	service     in.ExportSatellitesUseCase
	fileHandler FileHandler
}

func (controller ExportSatelliteController) Export(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filename := query.Get("filename")

	csvFile := controller.fileHandler.CreateFile(filename)
	defer controller.fileHandler.Close()

	err := controller.service.Export(csvFile)

	if err != nil {
		errorMessage := "Error while trying to export data, make sure that the query has the filename parameter"
		writeResponse(w, http.StatusInternalServerError, errorMessage)
	} else {
		writeResponse(w, http.StatusOK, "File has been written")
	}
}

func NewExportSatelliteController(service in.ExportSatellitesUseCase, fileHandler FileHandler) ExportSatelliteController {
	return ExportSatelliteController{
		service, fileHandler,
	}
}
