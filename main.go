package main

import (
	db "github.com/madebagus/ledgermate2/database" // Adjust the import path as needed
)

func main() {
	db.ConnectAndQueryPostgres()
	db.ConnectAndQueryMySQL()
}
