// Steam JP Finder
package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/mikoim/steam-jp-finder"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	// Redis
	redis := newPool("localhost:6379")
	defer redis.Close()

	// App
	a, err := newApp(redis, []byte("REPLACE BY YOUR STRONG KEY"))
	if err != nil {
		logrus.Fatal(err)
		return
	}

	// Routing
	r := mux.NewRouter()
	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/login", a.loginHandler)
	r.HandleFunc("/login/callback", a.loginCallbackHandler)
	r.HandleFunc("/logout", a.logoutHandler)

	// Logging
	h := sjf.Logging(r)

	logrus.Fatal(http.ListenAndServe(":8080", h))
}
