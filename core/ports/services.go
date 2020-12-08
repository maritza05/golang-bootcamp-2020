package ports

import (
	"github.com/maritza05/golang-bootcamp-2020/core/domain"
)

type SatelliteService interface {
	GetAll() ([]domain.Satellite, error)
}
