package main

import (
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/getbread/redistore"
	"github.com/gorilla/sessions"
	"github.com/mikoim/steam-jp-finder"
	"github.com/unrolled/render"
	"github.com/yohcop/openid-go"
)

const (
	openidURL   = "https://steamcommunity.com/openid"
	sessionName = "louise"
)

type app struct {
	pool           *redis.Pool
	rdr            *render.Render
	session        sessions.Store
	nonceStore     openid.NonceStore
	discoveryCache openid.DiscoveryCache
}

func newApp(pool *redis.Pool, keyPairs ...[]byte) (*app, error) {
	session, err := redistore.NewRediStoreWithPool(pool, keyPairs...)
	if err != nil {
		return nil, err
	}

	return &app{
		pool:           pool,
		rdr:            render.New(),
		session:        session,
		nonceStore:     openid.NewSimpleNonceStore(),
		discoveryCache: openid.NewSimpleDiscoveryCache(),
	}, nil
}

func (s *app) myHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := s.session.Get(r, sessionName)
	session.Values["foo"] = "bar"
	session.Values[42] = 43

	session.Save(r, w)

	w.Write([]byte("MyHandler"))
}

func (s *app) loginHandler(w http.ResponseWriter, r *http.Request) {
	if url, err := openid.RedirectURL(openidURL, sjf.RootURI(r)+"/login/callback", sjf.RootURI(r)); err == nil {
		http.Redirect(w, r, url, 303)
	} else {
		log.Print(err)
	}
}

func (s *app) loginCallbackHandler(w http.ResponseWriter, r *http.Request) {
	id, err := openid.Verify(sjf.URI(r), s.discoveryCache, s.nonceStore)
	if err == nil {
		log.Println(id)
	} else {
		log.Println(err)
	}
}
