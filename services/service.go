package services

import (
	"net/http"
)

type Service struct {
	Handler http.Handler
}
