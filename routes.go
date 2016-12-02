package main

import (
	"github.com/jpopesculian/wa-server/services"
	"net/http"
)

func NewRouter(config Config) *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/",
		services.NewStaticFileService(config.FileDir).Handler)

	router.Handle("/api/users/",
		services.NewLocalProxyService(3000).Handler)

	return router
}
