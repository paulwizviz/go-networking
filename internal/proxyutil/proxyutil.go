package proxyutil

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Redirector struct {
	targetURL *url.URL
}

func (r Redirector) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	req.URL.Host = r.targetURL.Host
	req.URL.Scheme = r.targetURL.Scheme
	req.Header.Set("X-Forward-Host", req.Header.Get("Host"))
	req.Host = r.targetURL.Host

	proxy := httputil.NewSingleHostReverseProxy(r.targetURL)
	proxy.ServeHTTP(res, req)
}
