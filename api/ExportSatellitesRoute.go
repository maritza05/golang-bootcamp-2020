package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/maritza05/golang-bootcamp-2020/satellites/adapters/repos"
	"github.com/maritza05/golang-bootcamp-2020/satellites/adapters/rest"
	"github.com/maritza05/golang-bootcamp-2020/satellites/app"
)

func RegisterExportSatellitesRoute(router *mux.Router, path string) {
	dbClient := getDbClient()
	postgresRepo := repos.NewPostgresRepo(dbClient)
	csvRepo, err := repos.NewCsvRepo(nil)
	if err != nil {
		panic("Error while initializing csv repository")
	}
	exportService := app.NewExportSatelliteService(postgresRepo, csvRepo)
	exportController := rest.NewExportSatelliteController(exportService, new(DefaultFileHandler))

	router.HandleFunc(path, exportController.Export).Methods(http.MethodPost)
}

type DefaultFileHandler struct {
	File *os.File
}

func (dh DefaultFileHandler) CreateFile(filename string) *os.File {
	csvFile, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	dh.File = csvFile
	return dh.File
}

func (dh DefaultFileHandler) Close() {
	dh.File.Close()
}

func getDbClient() *sqlx.DB {
	config := getCustomerRepositoryConfig()
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config["host"],
		config["port"],
		config["user"],
		config["password"],
		config["dbname"])
	client, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxIdleTime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func getCustomerRepositoryConfig() map[string]string {
	return map[string]string{
		"user":     os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASSWD"),
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
		"dbname":   os.Getenv("DB_NAME"),
	}
}
