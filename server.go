package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func NewServer(config Config) http.Server {
	var routes http.Handler = NewRouter(config)
	return http.Server{
		Handler: addMiddleware(routes),
		Addr:    fmt.Sprintf(":%d", config.Port),
	}
}

func addMiddleware(router http.Handler) http.Handler {
	handler := handlers.CompressHandler(router)
	handler = handlers.LoggingHandler(os.Stdout, handler)
	return handler
}
