package services

import (
	"fmt"
	"net/http/httputil"
	"net/url"
)

type ProxyService struct {
	Target  *url.URL
	Handler *httputil.ReverseProxy
}

func NewProxyService(target string) ProxyService {
	url, err := url.Parse(target)
	if err != nil {
		panic(err)
	}
	return ProxyService{
		Target:  url,
		Handler: httputil.NewSingleHostReverseProxy(url),
	}
}

func NewLocalProxyService(port int64) ProxyService {
	target := fmt.Sprintf("http://127.0.0.1:%d", port)
	return NewProxyService(target)
}
