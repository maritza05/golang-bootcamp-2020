package in

import (
	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"
)

type ListSatellitesUseCase interface {
	GetAll() ([]domain.Satellite, error)
}
