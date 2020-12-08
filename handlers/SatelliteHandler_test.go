package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maritza05/golang-bootcamp-2020/core/domain"
	"github.com/maritza05/golang-bootcamp-2020/core/services"
	repos "github.com/maritza05/golang-bootcamp-2020/repositories"
)

func Test_should_return_satellites_with_status_code_200(t *testing.T) {
	req := givenGetRequestTo("/satellites", t)

	satellites := givenSomeSatellites()

	repo := repos.NewMemRepo(satellites)
	service := services.NewDefaultSatelliteService(repo)
	handler := NewSatelliteHandler(service)

	rec := httptest.NewRecorder()

	handler.GetAll(rec, req)

	res := rec.Result()
	defer res.Body.Close()
}

func givenGetRequestTo(path string, t *testing.T) *http.Request {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Errorf("Could not create tests request")
	}
	return req
}

func givenSomeSatellites() []domain.Satellite {
	return []domain.Satellite{
		{OfficialName: "USA 139",
			Country:          "USA",
			Owner:            "National Reconnaissance Office (NRO)",
			Use:              "Military",
			Purpose:          "Earth Observation",
			LaunchDate:       "6/15/2016",
			ExpectedLifetime: "15",
			LaunchSite:       "Cape Canaveral",
			LaunchVehicle:    "Falcon 9",
		},
	}
}
