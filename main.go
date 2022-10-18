package main

import (
	"fmt"
	"log"
	"net/http"
	"travelit/router"

	_ "github.com/lib/pq"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
