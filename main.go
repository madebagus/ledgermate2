package main

import (
	db "github.com/madebagus/ledgermate2/database"
)

func main() {
	db.ConnectAndQueryPostgres()
	db.ConnectAndQueryMySQL()
}
