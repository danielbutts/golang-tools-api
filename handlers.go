package toolexchange

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ToolIndex(w http.ResponseWriter, r *http.Request) {
	DB_NAME := os.Getenv("DATABASE_URL")
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	rows, err := db.Query("SELECT id, name, image_url, is_borrowed, borrowed_on  FROM tools")

	var tools []Tool

	for rows.Next() {
		var ID int
		var Name string
		var ImageURL string
		var IsBorrowed bool
		var BorrowedOn time.Time
		err = rows.Scan(&ID, &Name, &ImageURL, &IsBorrowed, &BorrowedOn)
		checkErr(err)

		tool := Tool{ID, Name, ImageURL, IsBorrowed, BorrowedOn}
		tools = append(tools, tool)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tools); err != nil {
		panic(err)
	}
}

func ToolShow(w http.ResponseWriter, r *http.Request) {
	DB_NAME := os.Getenv("DATABASE_URL")
	vars := mux.Vars(r)
	ID := vars["Id"]
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	rows, err := db.Query("SELECT id, name, image_url, is_borrowed, borrowed_on FROM tools WHERE id = $1", ID)

	var tools []Tool

	for rows.Next() {
		var ID int
		var Name string
		var ImageURL string
		var IsBorrowed bool
		var BorrowedOn time.Time
		err = rows.Scan(&ID, &Name, &ImageURL, &IsBorrowed, &BorrowedOn)
		checkErr(err)

		tool := Tool{ID, Name, ImageURL, IsBorrowed, BorrowedOn}
		tools = append(tools, tool)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tools); err != nil {
		panic(err)
	}
}

func ToolUpdate(w http.ResponseWriter, r *http.Request) {
	DB_NAME := os.Getenv("DATABASE_URL")
	decoder := json.NewDecoder(r.Body)
	var tool Tool
	err := decoder.Decode(&tool)
	checkErr(err)
	defer r.Body.Close()

	vars := mux.Vars(r)
	ID := vars["Id"]
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	rows, err := db.Query("UPDATE tools SET name = $1, image_url = $2, is_borrowed = $3, borrowed_on = $4 WHERE id = $5", tool.Name, tool.ImageURL, tool.IsBorrowed, tool.BorrowedOn, ID)
	checkErr(err)

	rows, err = db.Query("SELECT id, name, image_url, is_borrowed, borrowed_on FROM tools WHERE id = $1", ID)
	checkErr(err)

	var tools []Tool

	for rows.Next() {
		var ID int
		var Name string
		var ImageURL string
		var IsBorrowed bool
		var BorrowedOn time.Time
		err = rows.Scan(&ID, &Name, &ImageURL, &IsBorrowed, &BorrowedOn)
		checkErr(err)

		tool := Tool{ID, Name, ImageURL, IsBorrowed, BorrowedOn}
		tools = append(tools, tool)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tools); err != nil {
		panic(err)
	}
}

func ToolInsert(w http.ResponseWriter, r *http.Request) {
	DB_NAME := os.Getenv("DATABASE_URL")

	decoder := json.NewDecoder(r.Body)
	var tool Tool
	err := decoder.Decode(&tool)
	checkErr(err)
	defer r.Body.Close()

	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	var lastInsertId int
	err = db.QueryRow("INSERT INTO tools (name, image_URL, is_borrowed, borrowed_on) values ($1, $2, $3, $4) returning id;", tool.Name, tool.ImageURL, tool.IsBorrowed, tool.BorrowedOn).Scan(&lastInsertId)
	checkErr(err)

	row := db.QueryRow("SELECT id, name, image_url, is_borrowed, borrowed_on FROM tools WHERE id = $1", lastInsertId)
	checkErr(err)

	var tools []Tool

	var ID int
	var Name string
	var ImageURL string
	var IsBorrowed bool
	var BorrowedOn time.Time
	err = row.Scan(&ID, &Name, &ImageURL, &IsBorrowed, &BorrowedOn)
	checkErr(err)

	tool = Tool{ID, Name, ImageURL, IsBorrowed, BorrowedOn}
	tools = append(tools, tool)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tools); err != nil {
		panic(err)
	}
}

func ToolDelete(w http.ResponseWriter, r *http.Request) {
	DB_NAME := os.Getenv("DATABASE_URL")

	vars := mux.Vars(r)
	ID := vars["Id"]

	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	var DeletedID int
	var Name string
	var ImageURL string
	var IsBorrowed bool
	var BorrowedOn time.Time
	row := db.QueryRow("DELETE FROM tools WHERE id = $1 RETURNING id, name, image_url, is_borrowed, borrowed_on;", ID)

	err = row.Scan(&ID, &Name, &ImageURL, &IsBorrowed, &BorrowedOn)
	checkErr(err)

	var tools []Tool

	tool := Tool{DeletedID, Name, ImageURL, IsBorrowed, BorrowedOn}
	tools = append(tools, tool)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tools); err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
