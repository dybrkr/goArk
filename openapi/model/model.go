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
	CharacterImage   string `json:"CharacterImage"`
	ExpeditionLevel  int    `json:"ExpeditionLevel"`
	PvpGradeName     string `json:"PvpGradeName"`
	TownLevel        int    `json:"TownLevel"`
	TownName         string `json:"TownName"`
	Title            string `json:"Title"`
	GuildMemberGrade string `json:"GuildMemberGrade"`
	GuildName        string `json:"GuildName"`
	Stats            []struct {
		Type    string   `json:"Type"`
		Value   string   `json:"Value"`
		Tooltip []string `json:"Tooltip"`
	} `json:"Stats"`
	Tendencies []struct {
		Type     string `json:"Type"`
		Point    int    `json:"Point"`
		MaxPoint int    `json:"MaxPoint"`
	} `json:"Tendencies"`
	ServerName         string `json:"ServerName"`
	CharacterName      string `json:"CharacterName"`
	CharacterLevel     int    `json:"CharacterLevel"`
	CharacterClassName string `json:"CharacterClassName"`
	ItemAvgLevel       string `json:"ItemAvgLevel"`
	ItemMaxLevel       string `json:"ItemMaxLevel"`
}

type ArmoryEquipment []struct {
	Type    string `json:"Type"`
	Name    string `json:"Name"`
	Icon    string `json:"Icon"`
	Grade   string `json:"Grade"`
	Tooltip string `json:"Tooltip"`
}

type ArmoryAvatar []struct {
	Type    string `json:"Type"`
	Name    string `json:"Name"`
	Icon    string `json:"Icon"`
	Grade   string `json:"Grade"`
	IsSet   bool   `json:"IsSet"`
	IsInner bool   `json:"IsInner"`
	Tooltip string `json:"Tooltip"`
}
