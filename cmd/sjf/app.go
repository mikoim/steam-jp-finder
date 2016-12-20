package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
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

func (s *app) indexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := s.session.Get(r, sessionName)
	s.rdr.Text(w, http.StatusOK, fmt.Sprintln(session.Values["SteamID"]))
}

func (s *app) loginHandler(w http.ResponseWriter, r *http.Request) {
	if url, err := openid.RedirectURL(openidURL, sjf.RootURI(r)+"/login/callback", sjf.RootURI(r)); err == nil {
		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		logrus.Error(err)
	}
}

func (s *app) loginCallbackHandler(w http.ResponseWriter, r *http.Request) {
	id, err := openid.Verify(sjf.URI(r), s.discoveryCache, s.nonceStore)
	if err != nil {
		logrus.Error(err)
		http.Redirect(w, r, sjf.RootURI(r)+"/?error=Login failed", http.StatusSeeOther)
		return
	}

	steamID, err := sjf.SteamId(id)
	if err != nil {
		logrus.Error(err)
		http.Redirect(w, r, sjf.RootURI(r)+"/?error=Invalid Steam ID", http.StatusSeeOther)
		return
	}

	session, _ := s.session.Get(r, sessionName)
	session.Values["SteamID"] = steamID
	session.Save(r, w)

	http.Redirect(w, r, sjf.RootURI(r), http.StatusSeeOther)
}

func (s *app) logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := s.session.Get(r, sessionName)
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, sjf.RootURI(r), http.StatusSeeOther)
}
