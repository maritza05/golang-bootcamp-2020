package repos_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/maritza05/golang-bootcamp-2020/satellites/adapters/repos"
	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"

	"github.com/stretchr/testify/assert"
)

func Test_can_initialize_csv_repository_with_valid_csv(t *testing.T) {
	satellite := givenASatellite()
	fakeCsv := givenValidCsv(satellite)

	repo, err := repos.NewCsvRepo(strings.NewReader(fakeCsv))

	assert.NotNil(t, repo)
	assert.Nil(t, err)
}

func Test_can_initialize_csv_repository_without_source(t *testing.T) {
	repo, err := repos.NewCsvRepo(nil)

	assert.NotNil(t, repo)
	assert.Nil(t, err)
}

func Test_cant_initialize_csv_repository_with_invalid_csv(t *testing.T) {
	fakeCsv := givenInvalidCsv()
	_, err := repos.NewCsvRepo(strings.NewReader(fakeCsv))

	assert.NotNil(t, err)
}

func Test_csv_repository_gets_satellites_from_valid_csv(t *testing.T) {
	satellite := givenASatellite()
	fakeCsv := givenValidCsv(satellite)

	repo, _ := repos.NewCsvRepo(strings.NewReader(fakeCsv))
	satellites, _ := repo.GetAll()

	assert.NotNil(t, repo)
	assert.Equal(t, len(satellites), 1)
	assert.IsType(t, satellites[0], satellite)
	assert.EqualValues(t, satellites[0], satellite)
}

func Test_returns_empty_satellites_on_csv_with_just_header(t *testing.T) {
	fakeCsv := givenCsvWithoutRecords()
	repo, _ := repos.NewCsvRepo(strings.NewReader(fakeCsv))

	satellites, _ := repo.GetAll()

	assert.Equal(t, len(satellites), 0)
}

func Test_writes_satellites_as_csv_records(t *testing.T) {
	var b bytes.Buffer
	repo, _ := repos.NewCsvRepo(nil)
	satellites := []domain.Satellite{givenASatellite()}
	err := repo.Write(satellites, &b)

	assert.Equal(t, givenValidCsv(satellites[0]), strings.TrimSpace(b.String()))
	assert.Nil(t, err)

}

func givenValidCsv(satellite domain.Satellite) string {
	header := getValidCsvHeader()
	row := strings.Join([]string{
		satellite.OfficialName,
		satellite.Country,
		satellite.Owner,
		satellite.Use,
		satellite.Purpose,
		satellite.LaunchDate,
		satellite.ExpectedLifetime,
		satellite.LaunchSite,
		satellite.LaunchVehicle}, ",")

	return fmt.Sprintf("%s\n%s", header, row)
}

func givenCsvWithoutRecords() string {
	return getValidCsvHeader()
}

func getValidCsvHeader() string {
	return "official_name,country,owner,uses,purpose,launch_date,expected_lifetime_years,launch_site,launch_vehicle"

}

func givenInvalidCsv() string {
	header := "name1,name2"
	row := "test1,test2"
	return fmt.Sprintf("%s\n%s", header, row)
}

func givenASatellite() domain.Satellite {
	return domain.Satellite{
		OfficialName:     "AAUSat-4",
		Country:          "NR",
		Owner:            "University of Aalborg",
		Use:              "Civil",
		Purpose:          "Earth Observation",
		LaunchDate:       "4/25/2016",
		ExpectedLifetime: "",
		LaunchSite:       "Guiana Space Center",
		LaunchVehicle:    "Soyuz 2.1a",
	}
}
