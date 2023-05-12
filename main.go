package main

import (
	"log"
	"net/http"

	router "github.com/Hasan-Iqtedar/crud-app-go/routes"
)

func main() {
	log.Println("Server listening on port 5050...")
	log.Fatal(http.ListenAndServe(":5050", router.Router()))
}
