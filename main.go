package main

import (
	"fmt"
	"github.com/jpopesculian/wa-server/config"
	"log"
)

func serve() {
	fmt.Printf("Serving on port %d\n", *config.Port)
	server := NewServer()
	log.Fatal(server.ListenAndServe())
}

func main() {
	config.Parse()
	serve()
}
