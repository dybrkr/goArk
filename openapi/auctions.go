package openapi

import (
	"github.com/dybrkr/goArk/openapi/model"
	"net/http"
)

const (
	GetAuctionOptions = "/auctions/options"
	SearchAuctionItem = "/auctions/items"
)

func (o *OpenAPI) GetAuctionOptions() (model.AuctionOption, error) {
	var resp model.AuctionOption

	err := SendRequest(o, http.MethodGet, GetAuctionOptions, nil, &resp)
	return resp, err
}

func (o *OpenAPI) SearchAuctionItem() (model.Auction, error) {
	var resp model.Auction

	// not tested
	req := model.RequestAuctionItems{
		ItemLevelMin:     0,
		ItemLevelMax:     0,
		ItemGradeQuality: 0,
		SkillOptions:     nil,
		EtcOptions:       nil,
		Sort:             "",
		CategoryCode:     0,
		CharacterClass:   "",
		ItemTier:         0,
		ItemGrade:        "",
		ItemName:         "",
		PageNo:           0,
		SortCondition:    "",
	}

	err := SendRequest(o, http.MethodPost, SearchAuctionItem, req, &resp)
	return resp, err
}
