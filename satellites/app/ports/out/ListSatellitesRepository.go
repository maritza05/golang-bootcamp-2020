package out

import (
	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"
)

type ListSatelliteRepository interface {
	GetAll() ([]domain.Satellite, error)
}
