package riot

type league struct {
	riot *Riot
}

type LeagueList struct {
	Tier     string       `json:"tier"`
	LeagueId string       `json:"leagueId"`
	Queue    string       `json:"queue"`
	Name     string       `json:"name"`
	Entries  []LeagueItem `json:"entries"`
}

type LeagueItem struct {
	SummonerId   string    `json:"summonerId"`
	SummonerName string    `json:"summonerName"`
	LeaguePoints uint      `json:"leaguePoints"`
	Rank         Rank      `json:"rank"`
	Wins         uint      `json:"wins"`
	Losses       uint      `json:"losses"`
	Veteran      bool      `json:"veteran"`
	Inactive     bool      `json:"inactive"`
	FreshBlood   bool      `json:"freshBlood"`
	HotStreak    bool      `json:"hotStreak"`
	MinSeries    MinSeries `json:"minSeries"`
}

type MinSeries struct {
	Losses   uint   `json:"losses"`
	Progress string `json:"progress"`
	Target   uint   `json:"target"`
	Wins     uint   `json:"wins"`
}

func (l *league) GetChallenger(queue Queue) (*LeagueList, error) {
	leagueList := new(LeagueList)
	if err := l.riot.client.Send("GET", "/lol/league/v4/challengerleagues/by-queue/"+string(queue), nil, leagueList); err != nil {
		return nil, err
	}
	return leagueList, nil
}
