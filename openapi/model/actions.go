package model

import "time"

type AuctionOption struct {
	MaxItemLevel       int           `json:"MaxItemLevel"`
	ItemGradeQualities []int         `json:"ItemGradeQualities"`
	SkillOptions       []SkillOption `json:"SkillOptions"`
	EtcOptions         []EtcOption   `json:"EtcOptions"`
	Categories         []Categories  `json:"Categories"`
	ItemGrades         []string      `json:"ItemGrades"`
	ItemTiers          []int         `json:"ItemTiers"`
	Classes            []string      `json:"Classes"`
}

type SkillOption struct {
	Value        int      `json:"Value"`
	Class        string   `json:"Class"`
	Text         string   `json:"Text"`
	IsSkillGroup bool     `json:"IsSkillGroup"`
	Tripods      []Tripod `json:"Tripods"`
}

type Tripod struct {
	Value int    `json:"Value"`
	Text  string `json:"Text"`
	IsGem bool   `json:"IsGem"`
}

type EtcOption struct {
	Value   int    `json:"Value"`
	Text    string `json:"Text"`
	EtcSubs []struct {
		Value int    `json:"Value"`
		Text  string `json:"Text"`
		Class string `json:"Class"`
	} `json:"EtcSubs"`
}

type Categories struct {
	Subs     []CategoryItem `json:"Subs"`
	Code     int            `json:"Code"`
	CodeName string         `json:"CodeName"`
}

type CategoryItem struct {
	Code     int    `json:"Code"`
	CodeName string `json:"CodeName"`
}

type RequestAuctionItems struct {
	ItemLevelMin     int                  `json:"ItemLevelMin"`
	ItemLevelMax     int                  `json:"ItemLevelMax"`
	ItemGradeQuality int                  `json:"ItemGradeQuality"`
	SkillOptions     []SearchDetailOption `json:"SkillOptions"`
	EtcOptions       []SearchDetailOption `json:"EtcOptions"`
	Sort             string               `json:"Sort"`
	CategoryCode     int                  `json:"CategoryCode"`
	CharacterClass   string               `json:"CharacterClass"`
	ItemTier         int                  `json:"ItemTier"`
	ItemGrade        string               `json:"ItemGrade"`
	ItemName         string               `json:"ItemName"`
	PageNo           int                  `json:"PageNo"`
	SortCondition    string               `json:"SortCondition"`
}

type SearchDetailOption struct {
	FirstOption  int `json:"FirstOption"`
	SecondOption int `json:"SecondOption"`
	MinValue     int `json:"MinValue"`
	MaxValue     int `json:"MaxValue"`
}

type Auction struct {
	PageNo     int           `json:"PageNo"`
	PageSize   int           `json:"PageSize"`
	TotalCount int           `json:"TotalCount"`
	Items      []AuctionItem `json:"Items"`
}

type AuctionItem struct {
	Name         string       `json:"Name"`
	Grade        string       `json:"Grade"`
	Tier         int          `json:"Tier"`
	Level        int          `json:"Level"`
	Icon         string       `json:"Icon"`
	GradeQuality int          `json:"GradeQuality"`
	AuctionInfo  AuctionInfo  `json:"AuctionInfo"`
	Options      []ItemOption `json:"Options"`
}

type AuctionInfo struct {
	StartPrice      int       `json:"StartPrice"`
	BuyPrice        int       `json:"BuyPrice"`
	BidPrice        int       `json:"BidPrice"`
	EndDate         time.Time `json:"EndDate"`
	BidCount        int       `json:"BidCount"`
	BidStartPrice   int       `json:"BidStartPrice"`
	IsCompetitive   bool      `json:"IsCompetitive"`
	TradeAllowCount int       `json:"TradeAllowCount"`
}

type ItemOption struct {
	Type             string `json:"Type"`
	OptionName       string `json:"OptionName"`
	OptionNameTripod string `json:"OptionNameTripod"`
	Value            int    `json:"Value"`
	IsPenalty        bool   `json:"IsPenalty"`
	ClassName        string `json:"ClassName"`
}
