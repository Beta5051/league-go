package lolgame

type liveClientData struct {
	lolGame *LolGame
}

type ActivePlayer struct {
	Abilities     Abilities `json:"abilities"`
	ChampionStats struct {
		AbilityHaste                 float64 `json:"abilityHaste"`
		AbilityPower                 float64 `json:"abilityPower"`
		Armor                        float64 `json:"armor"`
		ArmorPenetrationFlat         float64 `json:"armorPenetrationFlat"`
		ArmorPenetrationPercent      float64 `json:"armorPenetrationPercent"`
		AttackDamage                 float64 `json:"attackDamage"`
		AttackRange                  float64 `json:"attackRange"`
		AttackSpeed                  float64 `json:"attackSpeed"`
		BonusArmorPenetrationPercent float64 `json:"bonusArmorPenetrationPercent"`
		BonusMagicPenetrationPercent float64 `json:"bonusMagicPenetrationPercent"`
		CritChance                   float64 `json:"critChance"`
		CritDamage                   float64 `json:"critDamage"`
		CurrentHealth                float64 `json:"currentHealth"`
		HealShieldPower              float64 `json:"healShieldPower"`
		HealthRegenRate              float64 `json:"healthRegenRate"`
		LifeSteal                    float64 `json:"lifeSteal"`
		MagicLethality               float64 `json:"magicLethality"`
		MagicPenetrationFlat         float64 `json:"magicPenetrationFlat"`
		MagicPenetrationPercent      float64 `json:"magicPenetrationPercent"`
		MagicResist                  float64 `json:"magicResist"`
		MaxHealth                    float64 `json:"maxHealth"`
		MoveSpeed                    float64 `json:"moveSpeed"`
		Omnivamp                     float64 `json:"omnivamp"`
		PhysicalLethality            float64 `json:"physicalLethality"`
		PhysicalVamp                 float64 `json:"physicalVamp"`
		ResourceMax                  float64 `json:"resourceMax"`
		ResourceRegenRate            float64 `json:"resourceRegenRate"`
		ResourceType                 string  `json:"resourceType"`
		ResourceValue                float64 `json:"resourceValue"`
		SpellVamp                    float64 `json:"spellVamp"`
		Tenacity                     float64 `json:"tenacity"`
	} `json:"championStats"`
	CurrentGold        float64   `json:"currentGold"`
	FullRunes          FullRunes `json:"fullRunes"`
	Level              uint      `json:"level"`
	SummonerName       string    `json:"summonerName"`
	TeamRelativeColors bool      `json:"teamRelativeColors"`
}

type Abilities struct {
	E struct {
		AbilityLevel   uint   `json:"abilityLevel"`
		DisplayName    string `json:"displayName"`
		Id             string `json:"id"`
		RawDescription string `json:"rawDescription"`
		RawDisplayName string `json:"rawDisplayName"`
	} `json:"E"`
	Passive struct {
		DisplayName    string `json:"displayName"`
		Id             string `json:"id"`
		RawDescription string `json:"rawDescription"`
		RawDisplayName string `json:"rawDisplayName"`
	} `json:"Passive"`
	Q struct {
		AbilityLevel   uint   `json:"abilityLevel"`
		DisplayName    string `json:"displayName"`
		Id             string `json:"id"`
		RawDescription string `json:"rawDescription"`
		RawDisplayName string `json:"rawDisplayName"`
	} `json:"Q"`
	R struct {
		AbilityLevel   uint   `json:"abilityLevel"`
		DisplayName    string `json:"displayName"`
		Id             string `json:"id"`
		RawDescription string `json:"rawDescription"`
		RawDisplayName string `json:"rawDisplayName"`
	} `json:"R"`
	W struct {
		AbilityLevel   uint   `json:"abilityLevel"`
		DisplayName    string `json:"displayName"`
		Id             string `json:"id"`
		RawDescription string `json:"rawDescription"`
		RawDisplayName string `json:"rawDisplayName"`
	} `json:"W"`
}

type FullRunes struct {
	GeneralRunes []struct {
		DisplayName    string `json:"displayName"`
		Id             int    `json:"id"`
		RawDescription string `json:"rawDescription"`
		RawDisplayName string `json:"rawDisplayName"`
	} `json:"generalRunes"`
	Runes
	StatRunes []struct {
		Id             uint   `json:"id"`
		RawDescription string `json:"rawDescription"`
	} `json:"statRunes"`
}

type Player struct {
	ChampionName    string  `json:"championName"`
	IsBot           bool    `json:"isBot"`
	IsDead          bool    `json:"isDead"`
	Items           []Item  `json:"items"`
	Level           uint    `json:"level"`
	Position        string  `json:"position"`
	RawChampionName string  `json:"rawChampionName"`
	RespawnTimer    float64 `json:"respawnTimer"`
	Runes           Runes   `json:"runes"`
	Scores          Scores  `json:"scores"`
	SkinID          uint    `json:"skinID"`
	SummonerName    string  `json:"summonerName"`
	SummonerSpells  struct {
		SummonerSpellOne struct {
			DisplayName    string `json:"displayName"`
			RawDescription string `json:"rawDescription"`
			RawDisplayName string `json:"rawDisplayName"`
		} `json:"summonerSpellOne"`
		SummonerSpellTwo struct {
			DisplayName    string `json:"displayName"`
			RawDescription string `json:"rawDescription"`
			RawDisplayName string `json:"rawDisplayName"`
		} `json:"summonerSpellTwo"`
	} `json:"summonerSpells"`
	Team string `json:"team"`
}

type Item struct {
	CanUse         bool   `json:"canUse"`
	Consumable     bool   `json:"consumable"`
	Count          uint   `json:"count"`
	DisplayName    string `json:"displayName"`
	ItemId         uint   `json:"itemID"`
	Price          uint   `json:"price"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
	Slot           uint   `json:"slot"`
}

type Scores struct {
	Assists    uint    `json:"assists"`
	CreepScore uint    `json:"creepScore"`
	Deaths     uint    `json:"deaths"`
	Kills      uint    `json:"kills"`
	WardScore  float64 `json:"wardScore"`
}

type Runes struct {
	Keystone struct {
		DisplayName    string `json:"displayName"`
		Id             uint   `json:"id"`
		RawDescription string `json:"rawDescription"`
		RawDisplayName string `json:"rawDisplayName"`
	} `json:"keystone"`
	PrimaryRuneTree struct {
		DisplayName    string `json:"displayName"`
		Id             uint   `json:"id"`
		RawDescription string `json:"rawDescription"`
		RawDisplayName string `json:"rawDisplayName"`
	} `json:"primaryRuneTree"`
	SecondaryRuneTree struct {
		DisplayName    string `json:"displayName"`
		Id             uint   `json:"id"`
		RawDescription string `json:"rawDescription"`
		RawDisplayName string `json:"rawDisplayName"`
	} `json:"secondaryRuneTree"`
}

type Events struct {
	Events []struct {
		EventID   int     `json:"EventID"`
		EventName string  `json:"EventName"`
		EventTime float64 `json:"EventTime"`
	} `json:"Events"`
}

type GameData struct {
	GameMode   string  `json:"gameMode"`
	GameTime   float64 `json:"gameTime"`
	MapName    string  `json:"mapName"`
	MapNumber  uint    `json:"mapNumber"`
	MapTerrain string  `json:"mapTerrain"`
}

type AllGameData struct {
	ActivePlayer ActivePlayer `json:"activePlayer"`
	AllPlayers   []Player     `json:"allPlayers"`
	Events       Events       `json:"events"`
	GameData     GameData     `json:"gameData"`
}

func (lcd *liveClientData) get(path string, out any) error {
	return lcd.lolGame.client.Send("GET", "/liveclientdata"+path, nil, out)
}

func (lcd *liveClientData) GetAll() (*AllGameData, error) {
	allGameData := new(AllGameData)
	if err := lcd.get("/allgamedata", allGameData); err != nil {
		return nil, err
	}
	return allGameData, nil
}

func (lcd *liveClientData) GetActivePlayer() (*ActivePlayer, error) {
	activePlayer := new(ActivePlayer)
	if err := lcd.get("/activeplayer", activePlayer); err != nil {
		return nil, err
	}
	return activePlayer, nil
}

func (lcd *liveClientData) GetActivePlayerName() (string, error) {
	var name string
	if err := lcd.get("/activeplayername", &name); err != nil {
		return "", err
	}
	return name, nil
}

func (lcd *liveClientData) GetActivePlayerAbilities() (*Abilities, error) {
	abilities := new(Abilities)
	if err := lcd.get("/activeplayerabilities", abilities); err != nil {
		return nil, err
	}
	return abilities, nil
}

func (lcd *liveClientData) GetActivePlayerRunes() (*FullRunes, error) {
	runes := new(FullRunes)
	if err := lcd.get("/activeplayerrunes", runes); err != nil {
		return nil, err
	}
	return runes, nil
}

func (lcd *liveClientData) GetPlayers() ([]*Player, error) {
	var players []*Player
	if err := lcd.get("/playerlist", &players); err != nil {
		return nil, err
	}
	return players, nil
}

func (lcd *liveClientData) GetPlayerRunes(summonerName string) (*Runes, error) {
	runes := new(Runes)
	if err := lcd.get("/playermainrunes?summonerName="+summonerName, runes); err != nil {
		return nil, err
	}
	return runes, nil
}

func (lcd *liveClientData) GetPlayerItems(summonerName string) ([]*Item, error) {
	var items []*Item
	if err := lcd.get("/playeritems?summonerName="+summonerName, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (lcd *liveClientData) GetEvents() (*Events, error) {
	events := new(Events)
	if err := lcd.get("/eventdata", events); err != nil {
		return nil, err
	}
	return events, nil
}

func (lcd *liveClientData) GetGameData() (*GameData, error) {
	gameData := new(GameData)
	if err := lcd.get("/gamestats", gameData); err != nil {
		return nil, err
	}
	return gameData, nil
}
