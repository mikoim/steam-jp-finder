package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"remote": r.RemoteAddr,
			"method": r.Method,
			"path":   r.RequestURI,
			"host":   r.Host,
			"status": nil,
			"bytes":  nil,
		}).Debug("")
		h.ServeHTTP(w, r)
	})
}
