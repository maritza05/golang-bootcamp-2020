package services

import (
	"github.com/maritza05/golang-bootcamp-2020/core/domain"
	"github.com/maritza05/golang-bootcamp-2020/core/ports"
)

type DefaultSatelliteService struct {
	repo ports.SatelliteRepository
}

func (ds DefaultSatelliteService) GetAll() ([]domain.Satellite, error) {
	satellites, err := ds.repo.GetAll()
	if err != nil {
		return satellites, err
	}
	return nil, err
}

func NewDefaultSatelliteService(repo ports.SatelliteRepository) DefaultSatelliteService {
	return DefaultSatelliteService{repo}
}
