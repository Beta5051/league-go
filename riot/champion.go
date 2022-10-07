package riot

type champion struct {
	riot *Riot
}

type ChampionRotations struct {
	MaxNewPlayerLevel            uint   `json:"maxNewPlayerLevel"`
	FreeChampionIdsForNewPlayers []uint `json:"freeChampionIdsForNewPlayers"`
	FreeChampionIds              []uint `json:"freeChampionIds"`
}

func (c *champion) GetRotations() (*ChampionRotations, error) {
	rotations := new(ChampionRotations)
	if err := c.riot.client.Send("GET", "/lol/platform/v3/champion-rotations", nil, rotations); err != nil {
		return nil, err
	}
	return rotations, nil
}
