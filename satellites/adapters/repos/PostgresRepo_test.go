package repos_test

import (
	"errors"
	"testing"

	"github.com/maritza05/golang-bootcamp-2020/satellites/adapters/repos"
	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"

	"github.com/stretchr/testify/assert"
)

func Test_postgres_repo_returns_data_as_satellites(t *testing.T) {
	client := FakeSqlClient{}

	repo := repos.NewPostgresRepo(client)
	satellites, err := repo.GetAll()

	assert.Nil(t, err)
	assert.NotNil(t, satellites)
	assert.Equal(t, len(satellites), 1)
	assert.IsType(t, []domain.Satellite{}, satellites)
}

func Test_postgres_repo_returns_error_if_query_returns_error(t *testing.T) {
	client := FakeSqlClientWithError{}

	repo := repos.NewPostgresRepo(client)

	satellites, err := repo.GetAll()

	assert.NotNil(t, err)
	assert.Nil(t, satellites)

}

type FakeSqlClient struct {
}

func (f FakeSqlClient) Select(data interface{}, query string, more ...interface{}) error {
	satellite := repos.DBSatellite{
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
	val := (data).(*[]repos.DBSatellite)
	*val = []repos.DBSatellite{satellite}
	return nil
}

func (f FakeSqlClient) Close() error {
	return nil
}

type FakeSqlClientWithError struct {
}

func (f FakeSqlClientWithError) Select(data interface{}, query string, more ...interface{}) error {
	return errors.New("Something happened!")
}

func (f FakeSqlClientWithError) Close() error {
	return nil
}
