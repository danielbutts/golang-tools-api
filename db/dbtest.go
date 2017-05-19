package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	DB_NAME = "tools-rest-api-dev"
)

func main() {
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("Inserting Values")

	var lastInsertId int
	err = db.QueryRow("INSERT INTO tools(name, is_borrowed, borrowed_on) VALUES($1, $2, $3) returning id;", "hammer", "true", time.Now()).Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("last inserted id =", lastInsertId)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
