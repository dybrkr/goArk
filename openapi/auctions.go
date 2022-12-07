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

func (o *OpenAPI) SearchAuctionItem(req model.RequestMarketItems) (model.Auction, error) {
	var resp model.Auction

	err := SendRequest(o, http.MethodPost, SearchAuctionItem, req, &resp)
	return resp, err
}
