package database

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

func connect() {

	connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)

}
