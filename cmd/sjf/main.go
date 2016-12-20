// Steam JP Finder
package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/mikoim/steam-jp-finder"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	// App
	a, err := newApp(newPool("localhost:6379"), []byte("REPLACE BY YOUR STRONG KEY"))
	if err != nil {
		log.Fatal(err)
		return
	}

	// Routing
	r := mux.NewRouter()
	r.HandleFunc("/login", a.loginHandler)
	r.HandleFunc("/login/callback", a.loginCallbackHandler)
	r.HandleFunc("/set", a.myHandler)

	// Logging
	h := sjf.Logging(r)

	log.Fatal(http.ListenAndServe(":8080", h))
}
