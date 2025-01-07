package datastore

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func NewPostgresDBConn(host string, port int, user, password, dbName string) (*sqlx.DB, error) {
	datasource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)
	dbConn, err := sqlx.Open("postgres", datasource)

	if err != nil {
		logrus.Error("Connection Database Error Log: ", err)
		return nil, err
	}

	logrus.Info("Connection Database: ", host)
	logrus.Info("Connection ping to postgre: ", dbConn.Ping())

	return dbConn, nil
}
