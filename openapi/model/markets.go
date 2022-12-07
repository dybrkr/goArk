package model

type SortType string
type CategoryCode int
type CharacterClass string
type SortCondition string

const (
	GRADE             SortType = "GRADE"
	YDAY_AVG_PRICE    SortType = "YDAY_AVG_PRICE"
	RECENT_PRICE      SortType = "RECENT_PRICE"
	CURRENT_MIN_PRICE SortType = "CURRENT_MIN_PRICE"
)

const (
	ASC  SortCondition = "ASC"
	DESC SortCondition = "DESC"
)

type MarketOption struct {
	Categories []Category `json:"Categories"`
	ItemGrades []string   `json:"ItemGrades"`
	ItemTiers  []int      `json:"ItemTiers"`
	Classes    []string   `json:"Classes"`
}

type Category struct {
	Subs     []CategoryItem `json:"Subs"`
	Code     int            `json:"Code"`
	CodeName string         `json:"CodeName"`
}

// 일단은 auction에서 참조하는 것과 동일한듯 함, 만약 다를 경우 분리해야함
//type CategoryItem struct {
//	Code     int    `json:"Code"`
//	CodeName string `json:"CodeName"`
//}

type MarketItemStats struct {
	Name             string            `json:"Name"`
	TradeRemainCount int               `json:"TradeRemainCount,omitempty"`
	BundleCount      int               `json:"BundleCount"`
	Stats            []MarketStatsInfo `json:"Stats"`
}

type MarketStatsInfo struct {
	Date       string  `json:"Date"`
	AvgPrice   float64 `json:"AvgPrice"`
	TradeCount int     `json:"TradeCount"`
}

type RequestMarketItems struct {
	Sort           string `json:"Sort"`
	CategoryCode   int    `json:"CategoryCode"`
	CharacterClass string `json:"CharacterClass"`
	ItemTier       int    `json:"ItemTier"`
	ItemGrade      string `json:"ItemGrade"`
	ItemName       string `json:"ItemName"`
	PageNo         int    `json:"PageNo"`
	SortCondition  string `json:"SortCondition"`
}

type MarketList struct {
	PageNo     int          `json:"PageNo"`
	PageSize   int          `json:"PageSize"`
	TotalCount int          `json:"TotalCount"`
	Items      []MarketItem `json:"Items"`
}

type MarketItem struct {
	ID               int    `json:"Id"`
	Name             string `json:"Name"`
	Grade            string `json:"Grade"`
	Icon             string `json:"Icon"`
	BundleCount      int    `json:"BundleCount"`
	TradeRemainCount int    `json:"TradeRemainCount"`
	YDayAvgPrice     int    `json:"YDayAvgPrice"`
	RecentPrice      int    `json:"RecentPrice"`
	CurrentMinPrice  int    `json:"CurrentMinPrice"`
}
