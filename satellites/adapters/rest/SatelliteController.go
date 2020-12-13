package rest

import (
	"net/http"

	"github.com/maritza05/golang-bootcamp-2020/satellites/app/ports/in"
)

type ListSatellitesController struct {
	service in.ListSatellitesUseCase
}

func (controller ListSatellitesController) GetAll(w http.ResponseWriter, r *http.Request) {
	satellites, err := controller.service.GetAll()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err)
	} else {
		writeResponse(w, http.StatusOK, satellites)
	}
}

func NewListSatelliteController(service in.ListSatellitesUseCase) ListSatellitesController {
	return ListSatellitesController{
		service,
	}
}
