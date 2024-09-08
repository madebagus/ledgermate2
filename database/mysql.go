package db

import (
	"database/sql"
	"fmt"
	"log"

	config "conf" // Adjust the import path as needed

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func ConnectAndQueryMySQL() {
	dbConfig := config.LoadDBConfig().MySQL
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error connecting to the MySQL database:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM acl_list_new")
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}
