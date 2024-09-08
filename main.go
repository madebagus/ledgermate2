package main
	database "github.com/madebagus/ledgermate2/database"
func main() {
	database.ConnectAndQueryPostgres()
	database.ConnectAndQueryMySQL()
}
