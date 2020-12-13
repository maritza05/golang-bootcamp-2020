package rest_test

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maritza05/golang-bootcamp-2020/satellites/adapters/rest"
	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"
	"github.com/stretchr/testify/assert"
)

var satellites = []domain.Satellite{
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

type FakeListSatelliteService struct {
	satellites []domain.Satellite
}

func (s FakeListSatelliteService) GetAll() ([]domain.Satellite, error) {
	return s.satellites, nil
}

type FakeListSatelliteServiceWithError struct {
}

func (s FakeListSatelliteServiceWithError) GetAll() ([]domain.Satellite, error) {
	return nil, errors.New("Some internal error happened")
}

func Test_handler_should_return_200_status_code_with_valid_setup(t *testing.T) {
	service := FakeListSatelliteService{satellites}
	handler := rest.NewListSatelliteController(service)
	assert.HTTPStatusCode(t, handler.GetAll, "GET", "/satellites", nil, http.StatusOK)
}

func Test_handler_should_return_500_status_code_when_service_returns_error(t *testing.T) {
	handler := rest.NewListSatelliteController(FakeListSatelliteServiceWithError{})
	assert.HTTPStatusCode(t, handler.GetAll, "GET", "/satellites", nil, http.StatusInternalServerError)
}

func Test_should_return_json_satellites(t *testing.T) {
	service := FakeListSatelliteService{satellites}
	handler := rest.NewListSatelliteController(service)

	fakeRequest := givenGetRequest(t, "/satellites")

	rec := httptest.NewRecorder()

	handler.GetAll(rec, fakeRequest)

	fakeResponse := rec.Result()
	defer fakeResponse.Body.Close()

	assertHaveSameJsonData(t, fakeResponse.Body, satellites)

}

func assertHaveSameJsonData(t *testing.T, responseBody io.Reader, expected []domain.Satellite) {
	body, err := ioutil.ReadAll(responseBody)
	assert.Nil(t, err)

	var response []domain.Satellite
	err = json.Unmarshal(body, &response)

	assert.Nil(t, err)
	assert.Equal(t, len(response), len(expected))
}

func givenGetRequest(t *testing.T, path string) *http.Request {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Errorf("Could not create tests request")
	}
	return req
}
