package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	config "github.com/madebagus/ledgermate2/conf"

	// Mysql DB Driver
	_ "github.com/go-sql-driver/mysql"
)

// PSQLDBInit : init DB
func MYSQLDBInit() *sqlx.DB {
	dbConfig := config.LoadDBConfig().Postgres
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Host, dbConfig.Port)

	db, err := sqlx.Open("mysql", connStr)
	if err != nil {
		logrus.Fatal("MYSQL - DB open failed")
		return nil
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatal("MYSQL - DB ping failed", err)
		return nil
	}

	return db
}
