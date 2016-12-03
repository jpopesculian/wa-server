package services

import (
	"net/http"
)

type StaticFileService struct {
	Dir     http.Dir
	Handler http.Handler
}

func NewStaticFileService(dirString string) StaticFileService {
	dir := http.Dir(dirString)
	return StaticFileService{
		Dir:     dir,
		Handler: http.FileServer(dir),
	}
}
