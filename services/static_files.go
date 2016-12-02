package services

import (
	"net/http"
)

func NewStaticFileService(dir string) Service {
	return Service{
		Handler: http.FileServer(http.Dir(dir)),
	}
}
