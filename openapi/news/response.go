package news

type EventResponse []struct {
	Title      string `json:"Title"`
	Thumbnail  string `json:"Thumbnail"`
	Link       string `json:"Link"`
	StartDate  string `json:"StartDate"`
	EndDate    string `json:"EndDate"`
	RewardDate string `json:"RewardDate,omitempty"`
}
