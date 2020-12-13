package repos

import (
	"github.com/maritza05/golang-bootcamp-2020/satellites/domain"
)

type DBSatellite struct {
	OfficialName     string `db:"official_name"`
	Country          string `db:"country"`
	Owner            string `db:"owner"`
	Use              string `db:"use"`
	Purpose          string `db:"purpose"`
	LaunchDate       string `db:"launch_date"`
	ExpectedLifetime string `db:"expected_lifetime"`
	LaunchSite       string `db:"launch_site"`
	LaunchVehicle    string `db:"launch_vehicle"`
}

type ClientConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (s DBSatellite) toSatellite() domain.Satellite {
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

type SqlClient interface {
	Select(interface{}, string, ...interface{}) error
	Close() error
}

type PostgresRepo struct {
	client SqlClient
}

func (repo PostgresRepo) GetAll() ([]domain.Satellite, error) {
	sqlQuery := "SELECT * FROM satellites"
	dbsatellites := make([]DBSatellite, 0)
	err := repo.client.Select(&dbsatellites, sqlQuery)
	if err != nil {
		return nil, err
	}
	var satellites []domain.Satellite
	for _, s := range dbsatellites {
		satellites = append(satellites, s.toSatellite())
	}
	return satellites, nil
}

func (repo PostgresRepo) Close() {
	repo.client.Close()
}

func NewPostgresRepo(client SqlClient) PostgresRepo {
	return PostgresRepo{client}
}
