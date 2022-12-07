package openapi

import (
	"github.com/dybrkr/goArk/openapi/model"
	"github.com/dybrkr/goArk/request"
	"net/http"
	"net/url"
)

const (
	GetGuildRanking = "/guilds/rankings"
)

func (o *OpenAPI) GetGuildRanking(serverName string) ([]model.GuildRanking, error) {
	var resp []model.GuildRanking

	params := url.Values{}
	params.Add("serverName", serverName)
	subUrl := request.EncodeURLValues(GetGuildRanking, params)

	err := SendRequest(o, http.MethodGet, subUrl, nil, &resp)
	return resp, err
}
