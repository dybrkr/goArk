package openapi

import (
	"encoding/json"
	"fmt"
	"github.com/dybrkr/goArk/openapi/model"
	"net/http"
	"strings"
)

const (
	GetMarketOptions = "/markets/options"
	GetItemStat      = "/markets/items/%s"
	SearchItem       = "/markets/items"
)

type MarketResponse interface {
	model.MarketOption | []model.MarketItemStats | model.MarketList
}

func SendRequest[T MarketResponse](o *OpenAPI, method string, urlPath string, input interface{}, output *T) error {
	header := map[string]string{}
	header["accept"] = "application/json"
	header["authorization"] = "bearer " + o.AccessKey

	if input != nil && strings.EqualFold(method, http.MethodPost) {
		header["Content-Type"] = "application/json"
	}

	resp, err := o.SendRequest(method, urlPath, header, input)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(resp), output)
	if err != nil {
		return err
	}

	return nil
}

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

func (o *OpenAPI) SearchItem(sortType model.SortType,
	code model.CategoryCode,
	class model.CharacterClass,
	itemTier int,
	itemGrade string,
	itemName string,
	pageNo int,
	cond model.SortCondition) (model.MarketList, error) {

	var resp model.MarketList

	req := model.RequestMarketItems{
		Sort:           string(sortType),
		CategoryCode:   int(code),
		CharacterClass: string(class),
		ItemTier:       itemTier,
		ItemGrade:      itemGrade,
		ItemName:       itemName,
		PageNo:         pageNo,
		SortCondition:  string(cond),
	}

	err := SendRequest(o, http.MethodPost, SearchItem, req, &resp)

	return resp, err
}
