package sjf

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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

// SteamID extracts Steam ID from OpenID ID.
func SteamID(id string) (uint64, error) {
	p := strings.Split(id, "/")
	if len(p) != 6 {
		return 0, fmt.Errorf("invalid id %q", id)
	}
	return strconv.ParseUint(p[5], 10, 64)
}
