package repos

import (
	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"
)

type MemRepo struct {
	satellites []domain.Satellite
}

func (memRepo MemRepo) GetAll() ([]domain.Satellite, error) {
	return memRepo.satellites, nil
}

func NewMemRepo(satellites []domain.Satellite) MemRepo {
	return MemRepo{
		satellites,
	}
}
