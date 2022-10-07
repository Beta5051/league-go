package riot

type summoner struct {
	riot *Riot
}

type Summoner struct {
	AccountId     string `json:"accountId"`
	ProfileIconId uint   `json:"profileIconId"`
	RevisionDate  uint64 `json:"revisionDate"`
	Name          string `json:"name"`
	Id            string `json:"id"`
	Puuid         string `json:"puuid"`
	SummonerLevel uint64 `json:"summonerLevel"`
}

func (s *summoner) getBy(path string) (*Summoner, error) {
	summoner := new(Summoner)
	if err := s.riot.client.Send("GET", "/lol/summoner/v4/summoners"+path, nil, summoner); err != nil {
		return nil, err
	}
	return summoner, nil
}

func (s *summoner) GetByAccountId(id string) (*Summoner, error) {
	return s.getBy("/by-account/" + id)
}

func (s *summoner) GetByName(name string) (*Summoner, error) {
	return s.getBy("/by-name/" + name)
}

func (s *summoner) GetByPuuid(puuid string) (*Summoner, error) {
	return s.getBy("/by-puuid/" + puuid)
}

func (s *summoner) GetById(id string) (*Summoner, error) {
	return s.getBy("/" + id)
}

func (s *summoner) GetMe() (*Summoner, error) {
	return s.getBy("/me")
}
