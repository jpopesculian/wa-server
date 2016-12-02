package main

import (
	"fmt"
	"log"
)

func serve() {
	config := loadConfig()
	server := NewServer(config)

	fmt.Printf("Serving on port %d\n", config.Port)
	log.Fatal(server.ListenAndServe())
}

func main() {
	serve()
}
