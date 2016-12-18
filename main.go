// Steam JP Finder
package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/yohcop/openid-go"
)

const (
	openidURL   = "https://steamcommunity.com/openid"
	baseURL     = "http://localhost:8080"
	sessionName = "louise"
)

var (
	rdr            *render.Render
	store          sessions.Store
	nonceStore     openid.NonceStore
	discoveryCache openid.DiscoveryCache
)

func init() {
	store = sessions.NewCookieStore([]byte("REPLACE BY YOUR STRONG KEY"))
	nonceStore = openid.NewSimpleNonceStore()
	discoveryCache = openid.NewSimpleDiscoveryCache()

	log.SetLevel(log.DebugLevel)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionName)
	session.Values["foo"] = "bar"
	session.Values[42] = 43

	session.Save(r, w)

	w.Write([]byte("MyHandler"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if url, err := openid.RedirectURL(openidURL, baseURL+"/login/callback", baseURL); err == nil {
		http.Redirect(w, r, url, 303)
	} else {
		log.Print(err)
	}
}

func loginCallbackHandler(w http.ResponseWriter, r *http.Request) {
	id, err := openid.Verify(baseURL+r.URL.String(), discoveryCache, nonceStore)
	if err == nil {
		log.Println(id)
	} else {
		log.Println(err)
	}
}

func main() {
	// Render
	rdr = render.New()

	// Routing
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/login/callback", loginCallbackHandler)
	r.HandleFunc("/set", myHandler)

	// Logging
	h := Logging(r)

	log.Fatal(http.ListenAndServe(":8080", h))
}
