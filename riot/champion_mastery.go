package riot

import "fmt"

type championMastery struct {
	riot *Riot
}

type ChampionMastery struct {
	ChampionId                   uint64 `json:"championId"`
	ChampionLevel                uint   `json:"championLevel"`
	ChampionPoints               uint   `json:"championPoints"`
	LastPlayTime                 uint64 `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel uint64 `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel uint64 `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	TokensEarned                 uint   `json:"tokensEarned"`
	SummonerId                   string `json:"summonerId"`
}

func (c *championMastery) GetList(summonerId string) ([]*ChampionMastery, error) {
	var list []*ChampionMastery
	if err := c.riot.client.Send("GET", "/lol/champion-mastery/v4/champion-masteries/by-summoner/"+summonerId, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *championMastery) get(path string) (*ChampionMastery, error) {
	championMastery := new(ChampionMastery)
	if err := c.riot.client.Send("GET", "/lol/champion-mastery/v4/champion-masteries/by-summoner"+path, nil, championMastery); err != nil {
		return nil, err
	}
	return championMastery, nil
}

func (c *championMastery) GetByChampionId(summonerId string, championId uint64) (*ChampionMastery, error) {
	return c.get(fmt.Sprintf("/%s/by-champion/%d", summonerId, championId))
}

func (c *championMastery) GetTop(summonerId string) (*ChampionMastery, error) {
	return c.get("/" + summonerId + "/top")
}

func (c *championMastery) GetScore(summonerId string) (uint, error) {
	var out uint
	if err := c.riot.client.Send("GET", "/lol/champion-mastery/v4/scores/by-summoner/"+summonerId, nil, &out); err != nil {
		return 0, err
	}
	return out, nil
}
