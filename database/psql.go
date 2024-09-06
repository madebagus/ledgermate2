package database

import (
	"github.com/jmoiron/sqlx"

	"Golang1/conf"

	// PostgreSQL DB Driver
	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

// PSQLDBInit : init DB
func PSQLDBInit() *sqlx.DB {
	v, err := conf.GetConfig()
	dbname := v.GetString("datastore.psql.db")
	host := v.GetString("datastore.psql.host")
	user := v.GetString("datastore.psql.user")
	password := v.GetString("datastore.psql.password")
	port := v.GetString("datastore.psql.port")
	schema := v.GetString("datastore.psql.schema")
	maxcon := v.GetInt("datastore.psql.maxcon")
	maxidle := v.GetInt("datastore.psql.maxidle")

	db, err := sqlx.Open("postgres",
		"host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable search_path="+schema)
	if err != nil {
		logrus.Fatal("PQ - DB open failed")
		return nil
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatal("PQ - DB ping failed", err)
		return nil
	}

	db.SetMaxOpenConns(maxcon)
	db.SetMaxIdleConns(maxidle)
	return db
}
