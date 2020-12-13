package in

import "io"

type ExportSatellitesUseCase interface {
	Export(w io.Writer) error
}
