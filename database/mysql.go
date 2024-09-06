package database

import (
	"fmt"

	"github.com/madebagus/ledgermate2/conf"

	"github.com/jmoiron/sqlx"

	// Mysql DB Driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/sirupsen/logrus"
)

// MYSQLInit : init DB
func MYSQLInit() *sqlx.DB {
	v, err := conf.GetConfig()
	dbname := v.GetString("datastore.mysql.db")
	host := v.GetString("datastore.mysql.host")
	user := v.GetString("datastore.mysql.user")
	password := v.GetString("datastore.mysql.password")
	port := v.GetString("datastore.mysql.port")
	maxcon := v.GetInt("datastore.mysql.maxcon")
	maxidle := v.GetInt("datastore.mysql.maxidle")

	db, err := sqlx.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		logrus.Fatal("MYSQL DB open failed")
		fmt.Println(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatal("MYSQL DB ping failed 2 times")
		fmt.Println(err)
		return nil
	}

	db.SetMaxOpenConns(maxcon)
	db.SetMaxIdleConns(maxidle)

	return db
}
