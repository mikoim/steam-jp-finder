package main

import "net/http"

// RootURI re-constructs URI without path from *http.Request.
func RootURI(r *http.Request) string {
	protocol := "http"
	host := r.Host
	if r.TLS != nil {
		protocol = "https"
	}
	return protocol + "://" + host
}

// URI re-constructs URI from *http.Request.
func URI(r *http.Request) string {
	return RootURI(r) + r.RequestURI
}
