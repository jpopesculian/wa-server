package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

var router *http.ServeMux
var config map[string]interface{}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	router = http.NewServeMux()
	config = loadFlagsAndEnv(make(map[string]interface{}))

	router.HandleFunc("/", handler)

	server := handlers.CompressHandler(router)
	server = handlers.LoggingHandler(os.Stdout, server)

	address := fmt.Sprintf(":%d", config["port"])
	fmt.Printf("Serving at %s\n", address)
	log.Fatal(http.ListenAndServe(address, server))
}
