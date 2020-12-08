package handlers

import (
	"net/http"

	"github.com/maritza05/golang-bootcamp-2020/core/domain"
	"github.com/maritza05/golang-bootcamp-2020/core/ports"
)

type SatelliteHandler struct {
	service ports.SatelliteService
}

func (handler SatelliteHandler) GetAll(w http.ResponseWriter, r *http.Request) ([]domain.Satellite, error) {
	satellites, err := handler.service.GetAll()
	if err != nil {
		return satellites, nil
	}
	return nil, err
}

func NewSatelliteHandler(service ports.SatelliteService) SatelliteHandler {
	return SatelliteHandler{
		service,
	}
}
