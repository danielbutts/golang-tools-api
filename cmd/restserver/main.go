package main

import (
	"log"
	"net/http"

	"github.com/danielbutts/toolexchange"
)

func main() {

	router := toolexchange.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
