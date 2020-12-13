package repos

import (
	"io"

	"github.com/gocarina/gocsv"

	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"
)

type CsvSatellite struct {
	OfficialName     string `csv:"official_name"`
	Country          string `csv:"country"`
	Owner            string `csv:"owner"`
	Use              string `csv:"uses"`
	Purpose          string `csv:"purpose"`
	LaunchDate       string `csv:"launch_date"`
	ExpectedLifetime string `csv:"expected_lifetime_years"`
	LaunchSite       string `csv:"launch_site"`
	LaunchVehicle    string `csv:"launch_vehicle"`
}

func (s CsvSatellite) toSatellite() domain.Satellite {
	return domain.Satellite{
		OfficialName:     s.OfficialName,
		Country:          s.Country,
		Owner:            s.Owner,
		Use:              s.Use,
		Purpose:          s.Purpose,
		LaunchDate:       s.LaunchDate,
		ExpectedLifetime: s.ExpectedLifetime,
		LaunchSite:       s.LaunchSite,
		LaunchVehicle:    s.LaunchVehicle,
	}
}

type CsvRepo struct {
	satellites []domain.Satellite
}

func (r CsvRepo) GetAll() ([]domain.Satellite, error) {
	return r.satellites, nil
}

func (r CsvRepo) Write(satellites []domain.Satellite, w io.Writer) error {
	csvSatellites := parseToCsvSatellites(satellites)
	err := gocsv.Marshal(csvSatellites, w)
	if err != nil {
		return err
	}
	return nil
}

func NewCsvRepo(source io.Reader) (*CsvRepo, error) {
	if source == nil {
		return &CsvRepo{}, nil
	}
	satellites, err := getSatellitesFromCsv(source)
	if err != nil {
		return nil, err
	}
	return &CsvRepo{satellites}, nil
}

func getSatellitesFromCsv(handle io.Reader) ([]domain.Satellite, error) {
	satellites := []CsvSatellite{}

	gocsv.FailIfUnmatchedStructTags = true

	err := gocsv.Unmarshal(handle, &satellites)
	if err != nil {
		return nil, err
	}
	return parseToDomainSatellites(satellites), nil
}

func parseToDomainSatellites(csvSatellites []CsvSatellite) []domain.Satellite {
	var satellites []domain.Satellite
	for _, s := range csvSatellites {
		satellites = append(satellites, s.toSatellite())
	}
	return satellites
}

func parseToCsvSatellites(satellites []domain.Satellite) []CsvSatellite {
	var csvSatellites []CsvSatellite
	for _, s := range satellites {
		csvSatellites = append(csvSatellites, CsvSatellite{
			OfficialName:     s.OfficialName,
			Country:          s.Country,
			Owner:            s.Owner,
			Use:              s.Use,
			Purpose:          s.Purpose,
			LaunchDate:       s.LaunchDate,
			ExpectedLifetime: s.ExpectedLifetime,
			LaunchSite:       s.LaunchSite,
			LaunchVehicle:    s.LaunchVehicle,
		})
	}
	return csvSatellites
}
