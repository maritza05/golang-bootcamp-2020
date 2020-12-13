package rest_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/maritza05/golang-bootcamp-2020/satellites/adapters/rest"
	"github.com/stretchr/testify/assert"
)

type FakeExportSatelliteService struct {
}

func (spy FakeExportSatelliteService) Export(w io.Writer) error {
	return nil
}

type FakeExportSatelliteServiceWithError struct {
}

func (spy FakeExportSatelliteServiceWithError) Export(w io.Writer) error {
	return errors.New("Some error happened!")
}

type FakeFileHandler struct {
	filename string
}

func (spy *FakeFileHandler) CreateFile(filename string) *os.File {
	spy.filename = filename
	return nil
}

func (spy FakeFileHandler) Close() {
}

func Test_export_should_return_json_satellites(t *testing.T) {
	service := new(FakeExportSatelliteService)
	fileHandler := new(FakeFileHandler)
	controller := rest.NewExportSatelliteController(service, fileHandler)

	fakeRequest := givenPostRequest(t, "/export?filename=test.csv")

	rec := httptest.NewRecorder()

	controller.Export(rec, fakeRequest)

	fakeResponse := rec.Result()
	defer fakeResponse.Body.Close()

	assert.Equal(t, http.StatusOK, fakeResponse.StatusCode)
	assert.Equal(t, "test.csv", fileHandler.filename)
}

func Test_export_without_filename_should_return_error(t *testing.T) {
	service := new(FakeExportSatelliteServiceWithError)
	fileHandler := new(FakeFileHandler)
	controller := rest.NewExportSatelliteController(service, fileHandler)

	fakeRequest := givenPostRequest(t, "/export")

	rec := httptest.NewRecorder()

	controller.Export(rec, fakeRequest)

	fakeResponse := rec.Result()
	defer fakeResponse.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, fakeResponse.StatusCode)
	assert.Equal(t, "", fileHandler.filename)

}

func givenPostRequest(t *testing.T, path string) *http.Request {
	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		t.Errorf("Could not create tests request")
	}
	return req
}
