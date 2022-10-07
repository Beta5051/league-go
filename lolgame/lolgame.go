package lolgame

import (
	"encoding/json"
	"league/internal"
	"net/http"
)

type LolGame struct {
	client *internal.Client

	LiveClientData *liveClientData
}

func NewLolGame() *LolGame {
	lolGame := new(LolGame)
	lolGame.client = internal.NewClient(
		func() string {
			return "https://127.0.0.1:2999"
		},
		func(req *http.Request) {

		},
		func(resp *http.Response) error {
			out := Error{}
			if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
				return err
			}
			return out
		},
	)

	lolGame.LiveClientData = &liveClientData{lolGame: lolGame}
	return lolGame
}
