package app

import (
	"github.com/maritza05/golang-bootcamp-2020/satellites/app/ports/out"
	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"
)

type ListSatelliteService struct {
	repo out.ListSatelliteRepository
}

func (ds ListSatelliteService) GetAll() ([]domain.Satellite, error) {
	satellites, err := ds.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return satellites, err
}

func NewListSatelliteService(repo out.ListSatelliteRepository) ListSatelliteService {
	return ListSatelliteService{repo}
}
