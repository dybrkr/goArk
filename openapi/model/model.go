package model

type Event struct {
	Title      string `json:"Title"`
	Thumbnail  string `json:"Thumbnail"`
	Link       string `json:"Link"`
	StartDate  string `json:"StartDate"`
	EndDate    string `json:"EndDate"`
	RewardDate string `json:"RewardDate,omitempty"`
}

type CharacterInfo struct {
	ServerName         string `json:"ServerName"`
	CharacterName      string `json:"CharacterName"`
	CharacterLevel     int    `json:"CharacterLevel"`
	CharacterClassName string `json:"CharacterClassName"`
	ItemAvgLevel       string `json:"ItemAvgLevel"`
	ItemMaxLevel       string `json:"ItemMaxLevel"`
}

type ArmoryProfile struct {
	CharacterImage     string     `json:"CharacterImage"`
	ExpeditionLevel    int        `json:"ExpeditionLevel"`
	PvpGradeName       string     `json:"PvpGradeName"`
	TownLevel          int        `json:"TownLevel"`
	TownName           string     `json:"TownName"`
	Title              string     `json:"Title"`
	GuildMemberGrade   string     `json:"GuildMemberGrade"`
	GuildName          string     `json:"GuildName"`
	Stats              []Stat     `json:"Stats"`
	Tendencies         []Tendency `json:"Tendencies"`
	ServerName         string     `json:"ServerName"`
	CharacterName      string     `json:"CharacterName"`
	CharacterLevel     int        `json:"CharacterLevel"`
	CharacterClassName string     `json:"CharacterClassName"`
	ItemAvgLevel       string     `json:"ItemAvgLevel"`
	ItemMaxLevel       string     `json:"ItemMaxLevel"`
}

type Stat struct {
	Type    string   `json:"Type"`
	Value   string   `json:"Value"`
	Tooltip []string `json:"Tooltip"`
}

type Tendency struct {
	Type     string `json:"Type"`
	Point    int    `json:"Point"`
	MaxPoint int    `json:"MaxPoint"`
}

type ArmoryEquipment struct {
	Type    string `json:"Type"`
	Name    string `json:"Name"`
	Icon    string `json:"Icon"`
	Grade   string `json:"Grade"`
	Tooltip string `json:"Tooltip"`
}
type ArmoryAvatars []ArmoryAvatar
type ArmoryAvatar struct {
	Type    string `json:"Type"`
	Name    string `json:"Name"`
	Icon    string `json:"Icon"`
	Grade   string `json:"Grade"`
	IsSet   bool   `json:"IsSet"`
	IsInner bool   `json:"IsInner"`
	Tooltip string `json:"Tooltip"`
}

type ArmorySkill struct {
	Name        string        `json:"Name"`
	Icon        string        `json:"Icon"`
	Level       int           `json:"Level"`
	Type        string        `json:"Type"`
	IsAwakening bool          `json:"IsAwakening"`
	Tripods     []SkillTriPod `json:"Tripods"`
	Rune        SkillRune     `json:"Rune"`
	Tooltip     string        `json:"Tooltip"`
}

type SkillTriPod struct {
	Tier       int    `json:"Tier"`
	Slot       int    `json:"Slot"`
	Name       string `json:"Name"`
	Icon       string `json:"Icon"`
	Level      int    `json:"Level"`
	IsSelected bool   `json:"IsSelected"`
	Tooltip    string `json:"Tooltip"`
}

type SkillRune struct {
	Name    string `json:"Name"`
	Icon    string `json:"Icon"`
	Grade   string `json:"Grade"`
	Tooltip string `json:"Tooltip"`
}

type ArmoryEngraving struct {
	Engravings []Engraving `json:"Engravings"`
	Effects    []Effect    `json:"Effects"`
}

type Engraving struct {
	Slot    int    `json:"Slot"`
	Name    string `json:"Name"`
	Icon    string `json:"Icon"`
	Tooltip string `json:"Tooltip"`
}

type Effect struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type ArmoryCard struct {
	Cards   []Card       `json:"Cards"`
	Effects []CardEffect `json:"Effects"`
}

type Card struct {
	Slot       int    `json:"Slot"`
	Name       string `json:"Name"`
	Icon       string `json:"Icon"`
	AwakeCount int    `json:"AwakeCount"`
	AwakeTotal int    `json:"AwakeTotal"`
	Grade      string `json:"Grade"`
	Tooltip    string `json:"Tooltip"`
}

type CardEffect struct {
	Index     int      `json:"Index"`
	CardSlots []int    `json:"CardSlots"`
	Items     []Effect `json:"Items"`
}

type ArmoryGem struct {
	Gems    []Gem       `json:"Gems"`
	Effects []GemEffect `json:"Effects"`
}

type Gem struct {
	Slot    int    `json:"Slot"`
	Name    string `json:"Name"`
	Icon    string `json:"Icon"`
	Level   int    `json:"Level"`
	Grade   string `json:"Grade"`
	Tooltip string `json:"Tooltip"`
}

type GemEffect struct {
	GemSlot     int    `json:"GemSlot"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Icon        string `json:"Icon"`
	Tooltip     string `json:"Tooltip"`
}

type ColosseumInfo struct {
	Rank       int         `json:"Rank"`
	PreRank    int         `json:"PreRank"`
	Exp        int         `json:"Exp"`
	Colosseums []Colosseum `json:"Colosseums"`
}

type Colosseum struct {
	SeasonName      string                        `json:"SeasonName"`
	Competitive     AggregationTeamDeathMatchRank `json:"Competitive"`
	TeamDeathmatch  Aggregation                   `json:"TeamDeathmatch"`
	Deathmatch      Aggregation                   `json:"Deathmatch"`
	TeamElimination AggregationElimination        `json:"TeamElimination"`
	CoOpBattle      Aggregation                   `json:"CoOpBattle"`
}

type AggregationTeamDeathMatchRank struct {
	Rank         int    `json:"Rank"`
	RankName     string `json:"RankName"`
	RankIcon     string `json:"RankIcon"`
	RankLastMmr  int    `json:"RankLastMmr"`
	PlayCount    int    `json:"PlayCount"`
	VictoryCount int    `json:"VictoryCount"`
	LoseCount    int    `json:"LoseCount"`
	TieCount     int    `json:"TieCount"`
	KillCount    int    `json:"KillCount"`
	AceCount     int    `json:"AceCount"`
	DeathCount   int    `json:"DeathCount"`
}

type Aggregation struct {
	PlayCount    int `json:"PlayCount"`
	VictoryCount int `json:"VictoryCount"`
	LoseCount    int `json:"LoseCount"`
	TieCount     int `json:"TieCount"`
	KillCount    int `json:"KillCount"`
	AceCount     int `json:"AceCount"`
	DeathCount   int `json:"DeathCount"`
}

type AggregationElimination struct {
	FirstWinCount   int `json:"FirstWinCount"`
	SecondWinCount  int `json:"SecondWinCount"`
	ThirdWinCount   int `json:"ThirdWinCount"`
	FirstPlayCount  int `json:"FirstPlayCount"`
	SecondPlayCount int `json:"SecondPlayCount"`
	ThirdPlayCount  int `json:"ThirdPlayCount"`
	AllKillCount    int `json:"AllKillCount"`
	PlayCount       int `json:"PlayCount"`
	VictoryCount    int `json:"VictoryCount"`
	LoseCount       int `json:"LoseCount"`
	TieCount        int `json:"TieCount"`
	KillCount       int `json:"KillCount"`
	AceCount        int `json:"AceCount"`
	DeathCount      int `json:"DeathCount"`
}

type Collectible struct {
	Type              string             `json:"Type"`
	Icon              string             `json:"Icon"`
	Point             int                `json:"Point"`
	MaxPoint          int                `json:"MaxPoint"`
	CollectiblePoints []CollectiblePoint `json:"CollectiblePoints"`
}

type CollectiblePoint struct {
	PointName string `json:"PointName"`
	Point     int    `json:"Point"`
	MaxPoint  int    `json:"MaxPoint"`
}
