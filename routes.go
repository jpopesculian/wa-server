package main

import (
	"github.com/jpopesculian/wa-server/config"
	"github.com/jpopesculian/wa-server/services"
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/",
		services.NewStaticFileService(*config.StaticFileDir).Handler)

	router.Handle("/api/users/",
		services.NewProxyService(*config.UserServiceAddr).Handler)

	return router
}
