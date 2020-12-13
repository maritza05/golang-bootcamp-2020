package app

import (
	"io"

	"github.com/maritza05/golang-bootcamp-2020/satellites/app/ports/out"
)

type ExportSatelliteService struct {
	sourceRepo out.ListSatelliteRepository
	destRepo   out.ExportSatelliteRepository
}

func (es ExportSatelliteService) Export(w io.Writer) error {
	satellites, _ := es.sourceRepo.GetAll()
	err := es.destRepo.Write(satellites, w)
	if err != nil {
		return err
	}
	return nil
}

func NewExportSatelliteService(sourceRepo out.ListSatelliteRepository,
	destRepo out.ExportSatelliteRepository) ExportSatelliteService {
	return ExportSatelliteService{sourceRepo, destRepo}
}
