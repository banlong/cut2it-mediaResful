package main

import (
	"log"
	"net/http"

)

//Start service, listen on port 8080
func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
