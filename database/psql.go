package db

import (
	"database/sql"
	"fmt"
	"log"

	config "conf"

	// Adjust the import path as needed
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func ConnectAndQueryPostgres() {
	dbConfig := config.LoadDBConfig().Postgres
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Host, dbConfig.Port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the PostgreSQL database:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM hotel_list_new")
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
