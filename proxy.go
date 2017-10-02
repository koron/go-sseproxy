package sseproxy

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// Proxy is a HTTP reverse proxy which support SSE (Server-Sent Events).
type Proxy struct {
	*httputil.ReverseProxy
}

// New creates a SSE supported HTTP reverse proxy.
func New(u *url.URL) *Proxy {
	rp := httputil.NewSingleHostReverseProxy(u)
	rp.FlushInterval = 100 * time.Millisecond
	rp.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   24 * time.Hour,
			KeepAlive: 24 * time.Hour,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}
	return &Proxy{ReverseProxy: rp}
}
