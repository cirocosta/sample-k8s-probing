package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	address = flag.String("address", ":1313", "address to listen for requests")
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
}

func main() {
	http.HandleFunc("/liveness", logHandler)
	http.HandleFunc("/readiness", logHandler)

	fmt.Println("listening", *address)

	http.ListenAndServe(*address, nil)
}
