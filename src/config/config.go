package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/inventory_management")
	if err != nil {
		log.Fatal(err)
	}
}
