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

func main() {
	router = http.NewServeMux()
	config = loadFlagsAndEnv(make(map[string]interface{}))
	fs := http.FileServer(http.Dir("vendor/github.com/jpopesculian/wa-client/dist"))

	router.Handle("/", fs)

	server := handlers.CompressHandler(router)
	server = handlers.LoggingHandler(os.Stdout, server)

	address := fmt.Sprintf(":%d", config["port"])
	fmt.Printf("Serving at %s\n", address)
	log.Fatal(http.ListenAndServe(address, server))
}
