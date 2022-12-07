package openapi

import (
	"encoding/json"
	"github.com/dybrkr/goArk/openapi/model"
	"github.com/dybrkr/goArk/request"
	"net/http"
	"strings"
)

type OpenAPI struct {
	*request.RestAPI
	AccessKey string
}

func CreateAPI(endPoint string, accessKey string) *OpenAPI {
	return &OpenAPI{
		RestAPI:   request.CreateAPI(endPoint),
		AccessKey: accessKey,
	}
}

type APIResponse interface {
	// event
	[]model.Event |
		// character
		[]model.CharacterInfo |
		// auction
		model.AuctionOption | model.Auction |
		// market
		model.MarketOption | []model.MarketItemStats | model.MarketList |
		// guild
		[]model.GuildRanking |
		// armory
		model.ArmoryProfile | []model.ArmoryEquipment | model.ArmoryAvatars |
		[]model.ArmorySkill | model.ArmoryEngraving | model.ArmoryCard |
		model.ArmoryGem | model.ColosseumInfo | []model.Collectible
}

func SendRequest[T APIResponse](o *OpenAPI, method string, urlPath string, input interface{}, output *T) error {
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
