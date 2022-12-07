package openapi

import (
	"encoding/json"
	"github.com/dybrkr/goArk/openapi/model"
	"github.com/dybrkr/goArk/request"
	"net/http"
	"net/url"
)

const (
	GetGuildRanking = "/guilds/rankings"
)

type GuildResponse interface {
	[]model.GuildRanking
}

func SendGuildRequest[T GuildResponse](o *OpenAPI, urlPath string, input interface{}, output *T) error {
	header := map[string]string{}
	header["accept"] = "application/json"
	header["authorization"] = "bearer " + o.AccessKey

	if input != nil {
		// 인풋 있으면 처리해야하지만 일단 없고
	}

	resp, err := o.SendRequest(http.MethodGet, urlPath, header, nil)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(resp), output)
	if err != nil {
		return err
	}

	return nil
}

func (o *OpenAPI) GetGuildRanking(serverName string) ([]model.GuildRanking, error) {
	var resp []model.GuildRanking

	params := url.Values{}
	params.Add("serverName", serverName)

	subUrl := request.EncodeURLValues(GetGuildRanking, params)

	err := SendGuildRequest(o, subUrl, nil, &resp)
	return resp, err
}
