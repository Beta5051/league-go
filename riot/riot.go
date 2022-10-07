package riot

import (
	"encoding/json"
	"fmt"
	"league/internal"
	"net/http"
)

type Riot struct {
	ApiKey string
	Region Region
	client *internal.Client

	ChampionMastery *championMastery
	Champion        *champion
	League          *league
	Summoner        *summoner
}

func NewRiot(apiKey string, region Region) *Riot {
	riot := new(Riot)
	riot.ApiKey = apiKey
	riot.Region = region
	riot.client = internal.NewClient(
		func() string {
			return fmt.Sprintf("https://%s.api.riotgames.com", riot.Region)
		},
		func(req *http.Request) {
			req.Header.Add("X-Riot-Token", riot.ApiKey)
		},
		func(resp *http.Response) error {
			out := new(struct {
				Error Error `json:"status"`
			})
			if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
				return err
			}
			return out.Error
		},
	)

	riot.ChampionMastery = &championMastery{riot: riot}
	riot.Champion = &champion{riot: riot}
	riot.League = &league{riot: riot}
	riot.Summoner = &summoner{riot: riot}
	return riot
}
