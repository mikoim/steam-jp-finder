package main

import (
	"crypto/tls"
	"net/http"
	"testing"
)

func TestRootURI(t *testing.T) {
	var requests = []struct {
		in  *http.Request
		out string
	}{{
		&http.Request{Host: "example.com", RequestURI: "/foo/?bar=lol"},
		"http://example.com",
	}, {
		&http.Request{Host: "example.com", RequestURI: "/foo/?bar=lol", TLS: &tls.ConnectionState{}},
		"https://example.com",
	}}
	for i, r := range requests {
		o := RootURI(r.in)
		if o != r.out {
			t.Errorf("[%d] %q does not match %q", i, o, r.out)
		}
	}
}

func TestURI(t *testing.T) {
	var requests = []struct {
		in  *http.Request
		out string
	}{{
		&http.Request{Host: "example.com", RequestURI: "/foo/?bar=lol"},
		"http://example.com/foo/?bar=lol",
	}, {
		&http.Request{Host: "example.com", RequestURI: "/foo/?bar=lol", TLS: &tls.ConnectionState{}},
		"https://example.com/foo/?bar=lol",
	}}
	for i, r := range requests {
		o := URI(r.in)
		if o != r.out {
			t.Errorf("[%d] %q does not match %q", i, o, r.out)
		}
	}
}

func BenchmarkRootURI(b *testing.B) {
	dummy := &http.Request{
		Host:       "example.com",
		RequestURI: "/foo/?bar=lol",
	}
	for i := 0; i < b.N; i++ {
		RootURI(dummy)
	}
}

func BenchmarkURI(b *testing.B) {
	dummy := &http.Request{
		Host:       "example.com",
		RequestURI: "/foo/?bar=lol",
	}
	for i := 0; i < b.N; i++ {
		URI(dummy)
	}
}
