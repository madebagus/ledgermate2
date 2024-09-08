package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	config "github.com/madebagus/ledgermate2/conf"

	// Adjust the import path as needed
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// PSQLDBInit : init DB
func PSQLDBInit() *sqlx.DB {
	dbConfig := config.LoadDBConfig().Postgres
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Host, dbConfig.Port)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logrus.Fatal("PQ - DB open failed")
		return nil
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatal("PQ - DB ping failed", err)
		return nil
	}

	return db
}
