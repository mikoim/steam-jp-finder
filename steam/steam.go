// Package steam is wrapper for Steam Web API.
package steam

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

// Steam is Steam Web API client.
type Steam struct {
	request *gorequest.SuperAgent
	apiKey  string
}

// PlayerSummaries represents GetPlayerSummaries response.
type PlayerSummaries struct {
	Response playerSummariesResponse `json:"response"`
}

type playerSummariesResponse struct {
	Players []playerSummariesResponsePlayers `json:"players"`
}

type playerSummariesResponsePlayers struct {
	SteamID                  string `json:"steamid"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	ProfileState             int    `json:"profilestate"`
	PersonaName              string `json:"personaname"`
	LastLogoff               int    `json:"lastlogoff"`
	ProfileURL               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	PersonaState             int    `json:"personastate"`
	RealName                 string `json:"realname"`
	PrimaryClanID            string `json:"primaryclanid"`
	TimeCreated              int    `json:"timecreated"`
	PersonaStateFlags        int    `json:"personastateflags"`
	LocCountryCode           string `json:"loccountrycode"`
	LocStateCode             string `json:"locstatecode"`
	LocCityID                int    `json:"loccityid"`
}

// OwnedGames represents GetOwnedGames response.
type OwnedGames struct {
	Response ownedGamesResponse `json:"response"`
}

type ownedGamesResponse struct {
	GameCount int                       `json:"game_count"`
	Games     []ownedGamesResponseGames `json:"games"`
}

type ownedGamesResponseGames struct {
	AppID           int `json:"appid"`
	PlaytimeForever int `json:"playtime_forever"`
	Playtime2Weeks  int `json:"playtime_2weeks,omitempty"`
}

// NewSteam creates a new Steam Web API client.
func NewSteam(apiKey string) *Steam {
	return &Steam{
		request: gorequest.New(),
		apiKey:  apiKey,
	}
}

// ParsePlayerSummaries parses GetPlayerSummaries response
func ParsePlayerSummaries(resp *[]byte) (*PlayerSummaries, error) {
	p := PlayerSummaries{}
	if err := json.Unmarshal(*resp, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// ParseOwnedGames parses GetOwnedGames response
func ParseOwnedGames(resp *[]byte) (*OwnedGames, error) {
	o := OwnedGames{}
	if err := json.Unmarshal(*resp, &o); err != nil {
		return nil, err
	}
	return &o, nil
}
