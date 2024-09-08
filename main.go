package main

import (
	db "database" // Adjust the import path as needed
)

func main() {
	db.ConnectAndQueryPostgres()
	db.ConnectAndQueryMySQL()
}
