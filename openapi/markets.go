package openapi

import (
	"fmt"
	"github.com/dybrkr/goArk/openapi/model"
	"net/http"
)

const (
	GetMarketOptions = "/markets/options"
	GetItemStat      = "/markets/items/%s"
	SearchMarketItem = "/markets/items"
)

func (o *OpenAPI) GetMarketOptions() (model.MarketOption, error) {
	var resp model.MarketOption

	err := SendRequest(o, http.MethodGet, GetMarketOptions, nil, &resp)
	return resp, err
}

func (o *OpenAPI) GetItemStat(itemID string) ([]model.MarketItemStats, error) {
	var resp []model.MarketItemStats
	subUrl := fmt.Sprintf(GetItemStat, itemID)

	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}

func (o *OpenAPI) SearchItem(req model.RequestMarketItems) (model.MarketList, error) {

	var resp model.MarketList
	err := SendRequest(o, http.MethodPost, SearchMarketItem, req, &resp)

	return resp, err
}
