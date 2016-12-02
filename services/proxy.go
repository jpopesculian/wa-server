package services

import (
	"fmt"
	"net/http/httputil"
	"net/url"
)

type ProxyService struct {
	Handler *httputil.ReverseProxy
}

func NewProxyService(target string) ProxyService {
	url, _ := url.Parse(target)
	return ProxyService{
		Handler: httputil.NewSingleHostReverseProxy(url),
	}
}

func NewLocalProxyService(port int64) ProxyService {
	target := fmt.Sprintf("http://127.0.0.1:%d", port)
	return NewProxyService(target)
}
