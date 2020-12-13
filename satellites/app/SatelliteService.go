package app

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/maritza05/golang-bootcamp-2020/satellites/app/ports/out"
	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"
)

type DefaultSatelliteService struct {
	repo out.ListSatelliteRepository
}

func (ds DefaultSatelliteService) GetAll() ([]domain.Satellite, error) {
	satellites, err := ds.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return satellites, err
}

func (ds DefaultSatelliteService) Export(filename string) error {
	satellites, _ := ds.repo.GetAll()
	csvFile, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer csvFile.Close()

	err2 := gocsv.Marshal(&satellites, csvFile)
	if err2 != nil {
		return err2
	}
	return nil
}

func NewDefaultSatelliteService(repo out.ListSatelliteRepository) DefaultSatelliteService {
	return DefaultSatelliteService{repo}
}
