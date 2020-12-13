package out

import (
	"io"

	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"
)

type ExportSatelliteRepository interface {
	Write([]domain.Satellite, io.Writer) error
}
