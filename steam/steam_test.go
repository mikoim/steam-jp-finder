package steam

import (
	"reflect"
	"testing"
)

func TestNewSteam(t *testing.T) {
	apiKey := "0123456789ABCDEF0123456789ABCDEF"
	s := NewSteam(apiKey)
	if s.apiKey != apiKey {
		t.Errorf("Steam Web API key %q doesn't match %q.", s.apiKey, apiKey)
	}
	if s.request == nil {
		t.Error("HTTP client is nil.")
	}
}

func TestParsePlayerSummaries(t *testing.T) {
	var summaries = []struct {
		in  []byte
		out *PlayerSummaries
		err bool
	}{{
		[]byte(`{
  "response": {
    "players": [
      {
        "steamid": "76561197960435530",
        "communityvisibilitystate": 3,
        "profilestate": 1,
        "personaname": "Robin",
        "lastlogoff": 1482145710,
        "profileurl": "http://steamcommunity.com/id/robinwalker/",
        "avatar": "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/f1/f1dd60a188883caf82d0cbfccfe6aba0af1732d4.jpg",
        "avatarmedium": "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/f1/f1dd60a188883caf82d0cbfccfe6aba0af1732d4_medium.jpg",
        "avatarfull": "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/f1/f1dd60a188883caf82d0cbfccfe6aba0af1732d4_full.jpg",
        "personastate": 0,
        "realname": "Robin Walker",
        "primaryclanid": "103582791429521412",
        "timecreated": 1063407589,
        "personastateflags": 0,
        "loccountrycode": "US",
        "locstatecode": "WA",
        "loccityid": 3961
      }
    ]
  }
}`),
		&PlayerSummaries{
			playerSummariesResponse{
				[]playerSummariesResponsePlayers{{
					"76561197960435530",
					3,
					1,
					"Robin",
					1482145710,
					"http://steamcommunity.com/id/robinwalker/",
					"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/f1/f1dd60a188883caf82d0cbfccfe6aba0af1732d4.jpg",
					"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/f1/f1dd60a188883caf82d0cbfccfe6aba0af1732d4_medium.jpg",
					"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/f1/f1dd60a188883caf82d0cbfccfe6aba0af1732d4_full.jpg",
					0,
					"Robin Walker",
					"103582791429521412",
					1063407589,
					0,
					"US",
					"WA",
					3961,
				}},
			},
		},
		false,
	}, {
		[]byte("invalid json"),
		nil,
		true,
	}, {
		[]byte("{}"),
		&PlayerSummaries{},
		false,
	}}

	for i, s := range summaries {
		o, e := ParsePlayerSummaries(&s.in)
		if (e != nil) != s.err {
			t.Errorf("[%d] unexpected error %q", i, e)
		}
		if reflect.DeepEqual(o, s.out) == false {
			t.Errorf("[%d] %v does not match %v", i, o, s.out)
		}
	}
}

func TestParseOwnedGames(t *testing.T) {
	var ownedGames = []struct {
		in  []byte
		out *OwnedGames
		err bool
	}{{
		[]byte(`{
  "response": {
    "game_count": 2,
    "games": [
      {
        "appid": 10,
        "playtime_forever": 0
      },
      {
        "appid": 20,
        "playtime_forever": 0,
        "playtime_2weeks": 1
      }
    ]
  }
}`),
		&OwnedGames{
			ownedGamesResponse{
				2,
				[]ownedGamesResponseGames{{
					10,
					0,
					0,
				}, {
					20,
					0,
					1,
				}},
			},
		},
		false,
	}, {
		[]byte("invalid json"),
		nil,
		true,
	}, {
		[]byte("{}"),
		&OwnedGames{},
		false,
	}}

	for i, s := range ownedGames {
		o, e := ParseOwnedGames(&s.in)
		if (e != nil) != s.err {
			t.Errorf("[%d] unexpected error %q", i, e)
		}
		if reflect.DeepEqual(o, s.out) == false {
			t.Errorf("[%d] %v does not match %v", i, o, s.out)
		}
	}
}
