package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/DavidHuie/gomigrate"

	_ "github.com/lib/pq"
)

func main() {
	DB_NAME := os.Getenv("DATABASE_URL")
	var err error
	if len(os.Args) == 1 {
		err = errors.New("'up' or 'down' must be provided as an argument")
	}
	checkErr(err)
	arg := os.Args[1]

	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./")

	if arg == "up" {
		err = migrator.Migrate()
	} else if arg == "down" {
		err = migrator.Rollback()
	} else {
		err = errors.New(fmt.Sprintf("Invalid argument %s. Use 'up' or 'down'.", arg))
	}
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
